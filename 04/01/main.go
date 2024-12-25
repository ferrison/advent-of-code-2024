package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var w, h int

func main() {
    file, err := os.Open("../input-04")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var matrix []string
    for scanner.Scan() {
        line := scanner.Text()
        matrix = append(matrix, line)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    h = len(matrix)
    w = len(matrix[0])

    directions := [][]int{
        {-1, -1}, {-1,  0}, {-1,  1},
        { 0, -1},           { 0,  1},
        { 1, -1}, { 1,  0}, { 1,  1},
    }

    sum := 0
    for y, line := range matrix {
        for x := range line {
            for _, dir := range directions {
                dy := dir[0]
                dx := dir[1]
                if check(x, y, dx, dy, matrix) {
                    sum++
                }
            }

        }
    }
    fmt.Println(sum)
}

func check(x int, y int, dx int, dy int, matrix []string) bool {
    if string(matrix[y][x]) != "X" {
        return false
    }
    for _, expectedChar := range "MAS" {
        expectedChar := string(expectedChar)
        y = y+dy
        x = x+dx
        if y < 0 || y >= h || x < 0 || x >= w {
            return false
        }
        if string(matrix[y][x]) != expectedChar {
            return false
        }
    }
    return true
}
