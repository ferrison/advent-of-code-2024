package main

import (
	"bufio"
	"fmt"
	"log"
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
    BOX_LEFT = iota
    BOX_RIGHT = iota
    WALL = iota
    EMPTY = iota
    ROBOT = iota
)

var (
    NORTH = Vec{0, -1}
    EAST = Vec{1, 0}
    SOUTH = Vec{0, 1}
    WEST = Vec{-1, 0}
)

var grid [][]int
var moves []Vec
var robot Vec
var w, h int

func main() {
    file, err := os.Open("../input-15")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    grid = make([][]int, 0)
    y := 0
    for scanner.Scan() {
        line := scanner.Text()
        w = 2*len(line)

        if line == "" {
            break
        }

        row := make([]int, 0)
        for x, elR := range line {
            el := string(elR)
            switch el {
            case "#":
                row = append(row, WALL)
                row = append(row, WALL)
            case "O":
                row = append(row, BOX_LEFT)
                row = append(row, BOX_RIGHT)
            case "@":
                row = append(row, ROBOT)
                row = append(row, EMPTY)
                robot = Vec{x*2, y}
            case ".":
                row = append(row, EMPTY)
                row = append(row, EMPTY)
            }
        }
        grid = append(grid, row)
        y++
        h = y
    }

    moves = make([]Vec, 0)
    for scanner.Scan() {
        line := scanner.Text()
        
        for _, elR := range line {
            el := string(elR)
            switch el {
            case "^":
                moves = append(moves, NORTH)
            case ">":
                moves = append(moves, EAST)
            case "v":
                moves = append(moves, SOUTH)
            case "<":
                moves = append(moves, WEST)
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for _, move := range moves {
        moved := processMove(move, robot)
        if moved {
            robot = robot.Add(move)
        }
    }

    fmt.Println(calcSum(grid))

}

func processMove(move Vec, pos Vec) bool {
    currProp := grid[pos.y][pos.x]
    nextPos := pos.Add(move)
    nextProp := grid[nextPos.y][nextPos.x]
    if nextProp == WALL {
        return false
    }
    if nextProp == EMPTY {
        grid[nextPos.y][nextPos.x] = currProp
        grid[pos.y][pos.x] = EMPTY
        return true
    }
    if nextProp == BOX_LEFT || nextProp == BOX_RIGHT {
        if move == WEST || move == EAST {
            nextPropMoved := processMove(move, nextPos)
            if nextPropMoved {
                grid[nextPos.y][nextPos.x] = currProp
                grid[pos.y][pos.x] = EMPTY
                return true
            } else {
                return false
            }
        }

        currGrid := make([][]int, len(grid))
        for i := range grid {
            currGrid[i] = make([]int, len(grid[i]))
            copy(currGrid[i], grid[i])
        }

        var boxLeftMoved, boxRightMoved bool
        if nextProp == BOX_LEFT {
            boxLeftMoved = processMove(move, nextPos)
            boxRightMoved = processMove(move, nextPos.Add(Vec{1,0}))
        }
        if nextProp == BOX_RIGHT {
            boxRightMoved = processMove(move, nextPos)
            boxLeftMoved = processMove(move, nextPos.Add(Vec{-1,0}))
        }
        
        if boxLeftMoved && boxRightMoved {
            grid[nextPos.y][nextPos.x] = currProp
            grid[pos.y][pos.x] = EMPTY
            return true
        } else {
            grid = currGrid
            return false
        }
    }
    log.Fatal("Unexpected")
    return false
}

func printGrid(grid [][]int) {
    for _, row := range grid {
        for _, el := range row {
            switch el {
            case BOX_LEFT:
                fmt.Print("[")
            case BOX_RIGHT:
                fmt.Print("]")
            case WALL:
                fmt.Print("#")
            case ROBOT:
                fmt.Print("@")
            case EMPTY:
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}

func calcSum(grid [][]int) int {
    sum := 0
    for y, row := range grid {
        for x, el := range row {
            if el == BOX_LEFT {
                sum += y*100 + x
            }
        }
    }
    return sum
}
