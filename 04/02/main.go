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

    sum := 0
    for y, line := range matrix {
        for x := range line {
            if check(x, y, matrix) {
                sum++
            }
        }
    }
    fmt.Println(sum)
}

func check(x int, y int, matrix []string) bool {
    if string(matrix[y][x]) != "A" {
        return false
    }
    if y < 1 || y >= h-1 || x < 1 || x >= w-1 {
        return false
    }
    // check /
    corner1 := string(matrix[y+1][x-1])
    corner2 := string(matrix[y-1][x+1])
    if !((corner1 == "M" && corner2 == "S") || (corner1 == "S" && corner2 == "M")) {
        return false
    }

    // check \
    corner1 = string(matrix[y-1][x-1])
    corner2 = string(matrix[y+1][x+1])
    if !((corner1 == "M" && corner2 == "S") || (corner1 == "S" && corner2 == "M")) {
        return false
    }
    return true
}
