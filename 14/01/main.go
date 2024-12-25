package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func (v Vec) Mul(n int) Vec {
    return Vec{
        x: v.x*n,
        y: v.y*n,
    }
}

func main() {
    file, err := os.Open("../input-14")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    quadrants := make([]int, 4)
    re := regexp.MustCompile(`p=(\d*),(\d*) v=(-?\d*),(-?\d*)`)
    for scanner.Scan() {
        line := scanner.Text()

        parsedLine := re.FindStringSubmatch(line)
        px, _ := strconv.Atoi(parsedLine[1])
        py, _ := strconv.Atoi(parsedLine[2])
        vx, _ := strconv.Atoi(parsedLine[3])
        vy, _ := strconv.Atoi(parsedLine[4])

        pos := Vec{px, py}
        vel := Vec{vx, vy}

        newPos := process(pos, vel)

        if newPos.x < W/2 {
            if newPos.y < H/2 {
               quadrants[0]++ 
            }
            if newPos.y > H/2 {
                quadrants[2]++
            }
        }
        if newPos.x > W/2 {
            if newPos.y < H/2 {
               quadrants[1]++ 
            }
            if newPos.y > H/2 {
                quadrants[3]++
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    res := 1
    for _, q := range quadrants {
        res *= q
    }

    fmt.Println(res)

}

const (
    SECONDS = 100
    W = 101
    H = 103
)

func process(pos Vec, vel Vec) Vec {
    newPos := pos.Add(vel.Mul(SECONDS))
    newX := newPos.x%W
    newY := newPos.y%H
    if newX < 0 {
        newX = W + newX
    }
    if newY < 0 {
        newY = H + newY
    }
    return Vec{newX, newY}
}
