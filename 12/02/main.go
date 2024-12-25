//smell...
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

type Fence struct {
    x int
    y int
    direction Vec
}

var grid [][]string
var w, h int
var sum int

var processed [][]bool
var area int
var perimiter map[Fence]struct{}

func main() {
    file, err := os.Open("../input-12")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    grid = make([][]string, 0)
    for scanner.Scan() {
        line := scanner.Text()

        grid = append(grid, make([]string, 0))
        for _, charR := range line {
            char := string(charR)
            grid[len(grid)-1] = append(grid[len(grid)-1], char)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    h = len(grid)
    w = len(grid[0])

    processed = make([][]bool, h)
    for y := range grid {
        processed[y] = make([]bool, w)
    }

    sum = 0
    for y, row := range grid {
        for x, val := range row {
            area = 0
            perimiter = make(map[Fence]struct{})
            process(x, y, val)
            sides := processPerimiter()
            sum += area*sides
        }
    }

    fmt.Println(sum)
}


var NORTH = Vec{-1, 0}
var EAST = Vec{0, 1}
var SOUTH = Vec{1, 0}
var WEST = Vec{0, -1}

var side []Fence

func makeSide(fence Fence) {
    side = append(side, fence)
    delete(perimiter, fence)
    var directions []Vec
    if fence.direction == NORTH || fence.direction == SOUTH {
        directions = append(directions, []Vec{EAST, WEST}...)
    }
    if fence.direction == WEST || fence.direction == EAST {
        directions = append(directions, []Vec{NORTH, SOUTH}...)
    }
    for _, d := range directions {
        otherFence := Fence{fence.x+d.x, fence.y+d.y, fence.direction}
        if _, ok := perimiter[otherFence]; ok {
            makeSide(otherFence)
        }
    }

}

func processPerimiter() int {
    sides := make([][]Fence, 0)
    for len(perimiter) != 0 {
        side = make([]Fence, 0)
        var fence Fence
        for f := range perimiter {
            fence = f
            break
        }
        
        makeSide(fence)
        sides = append(sides, side)
    }
    return len(sides)
}

func process(x int, y int, val string) {
    if x < 0 || x >= w || y < 0 || y >= h {
        return
    }
    if processed[y][x] {
        return
    }
    if grid[y][x] != val {
        return
    }
    processed[y][x] = true
    
    area++
    for _, d := range []Vec{NORTH, EAST, SOUTH, WEST} {
        newX := x+d.x
        newY := y+d.y
        if !inbound(newX, newY) || val != grid[newY][newX] {
            perimiter[Fence{x, y, Vec{d.x, d.y}}] = struct{}{}
            continue
        }
        process(newX, newY, val)
    }
}

func inbound(x int, y int) bool {
    if x < 0 || x >= w || y < 0 || y >= h {
        return false
    }
    return true
}

