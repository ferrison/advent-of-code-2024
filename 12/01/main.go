package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var grid [][]string
var w, h int
var sum int

var processed [][]bool
var area, perimiter int

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
            perimiter = 0
            process(x, y, val)
            sum += area*perimiter
        }
    }

    fmt.Println(sum)
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
    for _, d := range [][]int{{-1,0}, {0, 1}, {1, 0}, {0, -1}} {
        dy := d[0]
        dx := d[1]
        newX := x+dx
        newY := y+dy
        if newX < 0 || newX >= w || newY < 0 || newY >= h {
            perimiter++
            continue
        }
        if val != grid[newY][newX] {
            perimiter++
        }
        process(newX, newY, val)
    }
}
