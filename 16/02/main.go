package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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

    scoresGrid := makeScoresGrid()
    places = make(map[Vec]struct{})
    countPlaces(scoresGrid, END)
    //printScores(scoresGrid)
    fmt.Println(len(places))
}

func process(currPos Vec, direction Vec, currScore int) (int, bool) {
    currProp := maze[currPos.y][currPos.x]
    if  currProp == WALL {
        return 0, false
    }
    if cachedScore, ok := scores[currPos]; ok {
        if cachedScore < currScore {
            return 0, false
        }
    }
    if currPos == END {
        scores[currPos] = currScore
        return currScore, true
    }
    
    scores[currPos] = currScore

    clockwiseTurn := Vec{direction.y, -direction.x}
    counterclockwiseTurn := Vec{-direction.y, direction.x}
    _, drOk := process(currPos.Add(direction), direction, currScore+1)
    _, cwOk := process(currPos.Add(clockwiseTurn), clockwiseTurn, currScore+1001)
    _, cntrCwOk := process(currPos.Add(counterclockwiseTurn), counterclockwiseTurn, currScore+1001)
    scrs := make([]int, 0)
    if drOk {
        scrs = append(scrs, currScore)
    }
    if cwOk {
        scrs = append(scrs, currScore+1000)
    }
    if cntrCwOk {
        scrs = append(scrs, currScore+1000)
    }
    if len(scrs) == 0 {
        return 0, false
    }
    scores[currPos] = slices.Min(scrs)
    return scores[currPos], true
}

var places map[Vec]struct{}

func countPlaces(scoresGrid [][]int, currPos Vec) {
    places[currPos] = struct{}{}
    currScore := scoresGrid[currPos.y][currPos.x]

    for _, d := range []Vec{NORTH, EAST, SOUTH, WEST} {
        nextPos := currPos.Add(d)
        nextScore := scoresGrid[nextPos.y][nextPos.x]
        if currScore - nextScore == 1001 || currScore - nextScore == 1 {
            countPlaces(scoresGrid, nextPos)
        }
    }
}

func makeScoresGrid() [][]int {
    scoresGrid := make([][]int, h)
    for i := range scoresGrid {
        row := make([]int, w)
        for j := range row {
            row[j] = math.MaxInt
        }
        scoresGrid[i] = row
    }
    for pos, score := range scores {
        scoresGrid[pos.y][pos.x] = score
    }
    return scoresGrid
}


func printScores(scoresGrid[][]int) {
    for _, row := range scoresGrid {
        for _, el := range row {
            if el == math.MaxInt {
                fmt.Print(".")
            } else {
                elStr := strconv.Itoa(el)
                fmt.Print("|"+elStr+"|")
            }
        }
        fmt.Println()
    }
}
