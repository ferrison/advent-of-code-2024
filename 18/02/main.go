package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vec struct {
    x int
    y int
}

func (v Vec) Add(other Vec) Vec {
    return Vec{
        x: v.x+other.x,
        y: v.y+other.y,
    }
}

const (
    W = 71
    H = 71
)

const (
    EMPTY = -1
    BYTE = -2
)

var bytesCount int

var grid [][]int

func main() {
    file, err := os.Open("../input-18")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    bytes := make([]Vec, 0)
    for scanner.Scan() {
        line := scanner.Text()

        coordsStr := strings.Split(line, ",")
        x, _ := strconv.Atoi(coordsStr[0])
        y, _ := strconv.Atoi(coordsStr[1])

        bytes = append(bytes, Vec{x, y})
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    lowBound := 0
    highBound := len(bytes)-1
    for {
        if lowBound == highBound {
            fmt.Println(bytesCount)
            fmt.Println(bytes[bytesCount])
            break
        }

        bytesCount = (highBound - lowBound)/2 + lowBound

        grid = makeGrid(bytes)
        calcShortest(Vec{0,0}, 0)

        distance := grid[H-1][W-1]
        if distance == EMPTY || distance == BYTE {
            highBound = bytesCount
        } else {
            lowBound = bytesCount + 1
        }
    }
}

func makeGrid(bytes []Vec) [][]int {
    grid := make([][]int, H)
    for y := range grid {
        row := make([]int, W)
        for x := range row {
            row[x] = EMPTY
        }
        grid[y] = row
    }
    for i:=0; i<bytesCount; i++ {
        byte := bytes[i]
        grid[byte.y][byte.x] = BYTE
    }
    return grid
}

func calcShortest(currPos Vec, currDistance int) {
    //fmt.Printf("x: %d\n", currPos.x)
    //fmt.Printf("y: %d\n", currPos.y)
    //fmt.Printf("currDistance: %d\n", currDistance)

    if currPos.x < 0 || currPos.x >= W || currPos.y < 0 || currPos.y >= H {
        return
    }
    currVal := grid[currPos.y][currPos.x]
    //fmt.Printf("currVal: %d\n", currVal)
    //fmt.Println()
    if currVal == BYTE {
        return
    }
    if currVal != EMPTY && currVal <= currDistance {
        return
    }
    grid[currPos.y][currPos.x] = currDistance
    for _, d := range []Vec{{0,-1}, {1, 0}, {0, 1}, {-1, 0}} {
        calcShortest(currPos.Add(d), currDistance+1)
    }
}

func printGrid(grid [][]int) {
    for _, row := range grid {
        for _, el := range row {
            switch el {
            case EMPTY:
                fmt.Print(".")
            case BYTE:
                fmt.Print("#")
            default:
                //fmt.Print(strconv.Itoa(el))
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}
