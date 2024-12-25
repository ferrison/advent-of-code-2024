package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
    EMPTY = -1
    WALL = -2
)

var grid [][]int
var w, h int
var start, end Vec

var cheats []int
var cheatDeltas []Vec

func main() {
    file, err := os.Open("../input-20")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    grid = make([][]int, 0)
    y := 0
    for scanner.Scan() {
        line := scanner.Text()

        row := make([]int, 0)
        for x, charR := range line {
            char := string(charR)

            switch char {
            case "#":
                row = append(row, WALL)
            case ".":
                row = append(row, EMPTY)
            case "S":
                start = Vec{x, y}
                row = append(row, EMPTY)
            case "E":
                end = Vec{x, y}
                row = append(row, EMPTY)
            }
        }
        grid = append(grid, row)
        y++
        w = len(line)
        h = y
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    findDistances(end, 0)
    //printGrid()

    cheats = make([]int, 0)
    cheatDeltas = generateAllCheatDeltas()
    findCheats(start)

    sum := 0
    for _, cheat := range cheats {
        if cheat >= 100 {
            sum++
        }
    }
    fmt.Println(sum)
}

func findDistances(currPos Vec, currDistance int) {
    currVal := grid[currPos.y][currPos.x]
    if currVal != EMPTY {
        return
    }
    grid[currPos.y][currPos.x] = currDistance
    for _, d := range []Vec{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
        findDistances(currPos.Add(d), currDistance+1)
    }
}

func findCheats(currPos Vec) {
    if currPos == end {
        return
    }
    currVal := grid[currPos.y][currPos.x]
    for _, d := range cheatDeltas {
        cheatPos := currPos.Add(d)
        if !inBound(cheatPos) {
            continue
        }
        cheatVal := grid[cheatPos.y][cheatPos.x]
        if cheatVal == WALL {
            continue
        }
        n := int(math.Abs(float64(d.x))) + int(math.Abs(float64(d.y)))
        if cheatVal < currVal - n {
            cheats = append(cheats, currVal - cheatVal - n)
        }
    }

    for _, d := range []Vec{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
        nextPos := currPos.Add(d)
        nextVal := grid[nextPos.y][nextPos.x]
        if nextVal == WALL {
            continue
        }
        if currVal - nextVal == 1 {
            findCheats(nextPos)
            break
        }
    }
}

func generateAllCheatDeltas() []Vec {
    deltas := make([]Vec, 0)
    for n:=2; n<=20; n++ {
        for x:=0; x<=n; x++ {
            y := n-x
            if x == 0 {
                deltas = append(deltas, Vec{x, y})
                deltas = append(deltas, Vec{x, -y})
                continue
            }
            if y == 0 {
                deltas = append(deltas, Vec{-x, y})
                deltas = append(deltas, Vec{x, y})
                continue
            }
            deltas = append(deltas, Vec{x, y})
            deltas = append(deltas, Vec{x, -y})
            deltas = append(deltas, Vec{-x, y})
            deltas = append(deltas, Vec{-x, -y})
        }
    }
    return deltas
}

func inBound(pos Vec) bool {
    if pos.x < 0 || pos.x >= w || pos.y < 0 || pos.y >= h {
        return false
    }
    return true
}


func printGrid() {
    for _, row := range grid {
        for _, el := range row {
            switch el {
            case WALL:
                fmt.Print("####")
            case EMPTY:
                fmt.Print("....")
            default:
                fmt.Printf("|%02d|", el)
            }
        }
        fmt.Println()
    }
}
