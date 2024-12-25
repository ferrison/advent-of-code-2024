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

type Robot struct {
    pos Vec
    vel Vec
}

var robots []Robot

func main() {
    file, err := os.Open("../input-14")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    robots = make([]Robot, 0)
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

        robots = append(robots, Robot{pos, vel})
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    var grid [][]string
    var curPoses map[Vec]struct{}

    printGrid(grid)

    secs:
    for i:=0; ; i++ {
        grid = make([][]string, H)
        curPoses = make(map[Vec]struct{})


        for _, robot := range robots {
            newPos := process(robot.pos, robot.vel, i)
            if _, ok := curPoses[newPos]; ok {
                continue secs
            }
            curPoses[newPos] = struct{}{}
        }

        for y:=0; y<H; y++ {
            row := make([]string, W)
            grid[y] = row
            for x:=0; x<W; x++ {
                row[x] = " "
            }
        }

        for robot := range curPoses {
            grid[robot.y][robot.x] = "*"
        }

        fmt.Println(i)
        printGrid(grid)
        fmt.Scanln()
    }

}

const (
    SECONDS = 100
    W = 101
    H = 103
)

func process(pos Vec, vel Vec, second int) Vec {
    newPos := pos.Add(vel.Mul(second))
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

func printGrid(grid [][]string) {
    out := ""
    for _, row := range grid {
        for _, el := range row {
            out += el
        }
        out += "\n"
    }
    fmt.Print(out)
}
