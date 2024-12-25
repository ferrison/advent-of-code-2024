package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
    file, err := os.Open("../input-21")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    sum := 0
    for scanner.Scan() {
        line := scanner.Text()

        presses := dirToNum(line)

        num, _ := strconv.Atoi(line[:len(line)-1])
        sum += dirToDir(presses, 0)*num

        fmt.Println(line)
        fmt.Println(dirToDir(presses, 0))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(sum)
}

type Vec struct {
    x int
    y int
}

func (v Vec) Sub(other Vec) Vec {
    return Vec{
        x: v.x-other.x,
        y: v.y-other.y,
    }
}

var numpad = map[string]Vec{
    "A": {2, 3},
    "0": {1, 3},
    "1": {0, 2},
    "2": {1, 2},
    "3": {2, 2},
    "4": {0, 1},
    "5": {1, 1},
    "6": {2, 1},
    "7": {0, 0},
    "8": {1, 0},
    "9": {2, 0},
}

func dirToNum(code string) string {
    dirpadPresses := ""
    currPos := numpad["A"]
    for _, bR := range code {
        b := string(bR)
        nextPos := numpad[b]
        move := nextPos.Sub(currPos)

        var horisontalPresses string
        if move.x >= 0 {
            horisontalPresses = strings.Repeat(">", move.x)
        } else {
            horisontalPresses = strings.Repeat("<", -move.x)
        }

        var verticalPresses string
        if move.y >= 0 {
            verticalPresses = strings.Repeat("v", move.y)
        } else {
            verticalPresses = strings.Repeat("^", -move.y)
        }

        presses := ""
        if currPos.y == 3 && nextPos.x == 0 {
            presses += verticalPresses + horisontalPresses
            dirpadPresses += presses + "A"
            currPos = nextPos
            continue
        }
        if currPos.x == 0 && nextPos.y == 3 {
            presses += horisontalPresses + verticalPresses
            dirpadPresses += presses + "A"
            currPos = nextPos
            continue
        }
        if currPos.x > nextPos.x {
            presses += horisontalPresses
        }
        if currPos.y > nextPos.y {
            presses += verticalPresses
        }
        if currPos.y < nextPos.y {
            presses += verticalPresses
        }
        if currPos.x < nextPos.x {
            presses += horisontalPresses
        }

        dirpadPresses += presses + "A"
        currPos = nextPos

    }
    return dirpadPresses
}

var dirpad = map[string]Vec{
    "A": {2, 0},
    "^": {1, 0},
    "<": {0, 1},
    "v": {1, 1},
    ">": {2, 1},
}

type Args struct {
    code string
    level int
}

var cache = map[Args]int{}

func dirToDir(code string, level int) int {
    if res, ok := cache[Args{code, level}]; ok {
        return res
    }
    if level == 25 {
        return len(code)
    }
    currPos := dirpad["A"]

    sum := 0
    for _, bR := range code {
        b := string(bR)
        nextPos := dirpad[b]
        move := nextPos.Sub(currPos)

        var horisontalPresses string
        if move.x >= 0 {
            horisontalPresses = strings.Repeat(">", move.x)
        } else {
            horisontalPresses = strings.Repeat("<", -move.x)
        }

        var verticalPresses string
        if move.y >= 0 {
            verticalPresses = strings.Repeat("v", move.y)
        } else {
            verticalPresses = strings.Repeat("^", -move.y)
        }

        presses := ""
        if currPos.y == 0 && nextPos.x == 0 {
            presses = verticalPresses + horisontalPresses
            sum += dirToDir(presses + "A", level+1)
            currPos = nextPos
            continue
        }
        if currPos.x == 0 && nextPos.y == 0 {
            presses = horisontalPresses + verticalPresses
            sum += dirToDir(presses + "A", level+1)
            currPos = nextPos
            continue
        }
        if currPos.x > nextPos.x {
            presses += horisontalPresses
        }
        if currPos.y > nextPos.y {
            presses += verticalPresses
        }
        if currPos.y < nextPos.y {
            presses += verticalPresses
        }
        if currPos.x < nextPos.x {
            presses += horisontalPresses
        }

        sum += dirToDir(presses + "A", level+1)
        currPos = nextPos
    }

    cache[Args{code, level}] = sum
    return sum
}
