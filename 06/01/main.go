package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
    x int
    y int
}

func (p Point) Add(other Point) Point {
    return Point{
        x: p.x + other.x,
        y: p.y + other.y,
    }
}

func main() {
    file, err := os.Open("../input-06")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var w, h int
    y := 0
    obstacles := make(map[Point]struct{})
    var guard Point
    for scanner.Scan() {
        line := scanner.Text()
        w = len(line)

        for x, charB := range line {
            char := string(charB)

            switch char {
            case "#":
                obstacles[Point{x, y}] = struct{}{}
            case "^":
                guard = Point{x, y}
            }
        }
        y++
        h = y
    }

    visited := make(map[Point]struct{})
    direction := Point{0, -1}
    for true {
        if guard.x < 0 || guard.y < 0 || guard.x >= w || guard.y >= h {
            break
        }

        visited[guard] = struct{}{}
        
        nextPos := guard.Add(direction)
        if _, ok := obstacles[nextPos]; ok {
            switch direction {
            case Point{0, -1}:  //North
                direction = Point{1, 0}  //East
            case Point{1, 0}:  //East
                direction = Point{0, 1}  //South
            case Point{0, 1}:  //South
                direction = Point{-1, 0}  //West
            case Point{-1, 0}:  //South
                direction = Point{0, -1}  //West
            }
            continue
        }

        guard = nextPos
    }



    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("w: %d\n", w)
    fmt.Printf("h: %d\n", h)
    fmt.Printf("Distinct visited: %d\n", len(visited))
}
