package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "github.com/schollz/progressbar/v3"
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

type Vector struct {
    point Point
    direction Point
}

var obstacles map[Point]struct{}
var w, h int

func main() {
    file, err := os.Open("../input-06")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    y := 0
    obstacles = make(map[Point]struct{})
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

    sum := 0
    bar := progressbar.Default(int64(h*w))
    for x:=0; x<=w; x++ {
        for y:=0; y<=h; y++ {
            bar.Add(1)
            newObstacle := Point{x, y}
            if _, ok := obstacles[newObstacle]; ok {
                continue
            }
            if newObstacle == guard {
                continue
            }
            
            if check(guard, newObstacle) {
                sum++
            }
        }
    }


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println()
    fmt.Printf("w: %d\n", w)
    fmt.Printf("h: %d\n", h)
    fmt.Println(sum)
}

func check(guard Point, newObstacle Point) bool {
    visited := make(map[Vector]struct{})
    direction := Point{0, -1}
    for true {
        if guard.x < 0 || guard.y < 0 || guard.x >= w || guard.y >= h {
            return false
        }

        if _, ok := visited[Vector{guard, direction}]; ok {
            return true
        }

        visited[Vector{guard, direction}] = struct{}{}
        
        nextPos := guard.Add(direction)
        if _, ok := obstacles[nextPos]; ok || nextPos == newObstacle {
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
    return false
}
