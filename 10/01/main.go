package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Point struct {
    x int
    y int
}

var grid [][]int
var totalNines map[Point]struct{}
var w, h int

func main() {
    file, err := os.Open("../input-10")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    grid = make([][]int, 0)
    for scanner.Scan() {
        line := scanner.Text()

        grid = append(grid, make([]int, 0))
        for _, charR := range line {
            val, _ := strconv.Atoi(string(charR))
            grid[len(grid)-1] = append(grid[len(grid)-1], val)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    w = len(grid[0])
    h = len(grid)

    score := 0
    for y, row := range grid {
        for x := range row {
            totalNines = make(map[Point]struct{})
            check(x, y, 0)
            score += len(totalNines)
        }
    }

    fmt.Println(score)
}

func check(x int, y int, curr int) {
    if x < 0 || x >= w || y < 0 || y >= h {
        return
    }
    if curr==9 && grid[y][x] == 9 {
        totalNines[Point{x, y}] = struct{}{}
        return
    }
    if grid[y][x] != curr {
        return
    }
    for _, d := range [][]int{{-1,0}, {0, 1}, {1, 0}, {0, -1}} {
        dx := d[0]
        dy := d[1]
        check(x+dx, y+dy, curr+1)
    }
}
