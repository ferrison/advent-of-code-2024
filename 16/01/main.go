package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

var (
    NORTH = Vec{0, -1}
    EAST = Vec{1, 0}
    SOUTH = Vec{0, 1}
    WEST = Vec{-1, 0}
)

const (
    WALL = iota
    EMPTY = iota
)

var maze [][]int
var scores map[Vec]int
var w, h int

var START, END Vec

func main() {
    file, err := os.Open("../input-16")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    maze = make([][]int, 0)
    y := 0
    for scanner.Scan() {
        line := scanner.Text()
        w = len(line)
        row := make([]int, 0)
        for x, charR := range line {
            char := string(charR)
            switch char {
            case "#":
                row = append(row, WALL)
            case ".":
                row = append(row, EMPTY)
            case "S":
                row = append(row, EMPTY)
                START = Vec{x, y}
            case "E":
                row = append(row, EMPTY)
                END = Vec{x, y}
            }
        }
        maze = append(maze, row)
        y++
        h = y
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    scores = make(map[Vec]int)
    process(START, EAST, 0)

    //printScores()
    fmt.Println(scores[END])
}

func process(currPos Vec, direction Vec, currScore int) {
    currProp := maze[currPos.y][currPos.x]
    if  currProp == WALL {
        return
    }
    if cachedScore, ok := scores[currPos]; ok {
        if cachedScore < currScore {
            return
        }
    }
    scores[currPos] = currScore
    if currPos == END {
        return
    }

    clockwiseTurn := Vec{direction.y, -direction.x}
    counterclockwiseTurn := Vec{-direction.y, direction.x}
    process(currPos.Add(direction), direction, currScore+1)
    process(currPos.Add(clockwiseTurn), clockwiseTurn, currScore+1001)
    process(currPos.Add(counterclockwiseTurn), counterclockwiseTurn, currScore+1001)
}

func printScores() {
    scoresGrid := make([][]string, h)
    for i := range scoresGrid {
        row := make([]string, w)
        for j := range row {
            row[j] = "."
        }
        scoresGrid[i] = row
    }
    for pos, score := range scores {
        scoresGrid[pos.y][pos.x] = "|"+strconv.Itoa(score)+"|"
    }
    for _, row := range scoresGrid {
        for _, el := range row {
            fmt.Print(el)
        }
        fmt.Println()
    }
}
