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

        presses := dirToDir(dirToDir(dirToNum(line)))
        fmt.Println(line)
        fmt.Println(dirToNum(line))
        fmt.Println(dirToDir(dirToNum(line)))
        fmt.Println(presses)
        fmt.Println(len(presses))

        num, _ := strconv.Atoi(line[:len(line)-1])
        sum += len(presses)*num
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
        if currPos.y < nextPos.y {
            presses += verticalPresses
        }
        if currPos.x < nextPos.x {
            presses += horisontalPresses
        }
        if currPos.y > nextPos.y {
            presses += verticalPresses
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
func dirToDir(code string) string {
    dirpadPresses := ""
    currPos := dirpad["A"]
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

        var presses string
        if move.y >= 0 {
            presses = strings.Repeat("v", move.y) + horisontalPresses
        } else {
            presses = horisontalPresses + strings.Repeat("^", -move.y)
        }
        dirpadPresses += presses + "A"

        currPos = nextPos
    }
    return dirpadPresses
}
