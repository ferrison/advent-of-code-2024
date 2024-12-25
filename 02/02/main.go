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
    file, err := os.Open("../input-02")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    safe := 0
    for scanner.Scan() {
        line := scanner.Text()

        levelStrs := strings.Fields(line)
        var levels []int

        for _, levelStr := range levelStrs {
            level, err := strconv.Atoi(levelStr)
            if err != nil {
                log.Fatal(err)
            }
            levels = append(levels, level)
        }

        if isSafe(levels) {
            safe++
            continue
        }

        safeFound := false
        for i:=0; i<len(levels); i++ {
            dLevel := make([]int, 0, len(levels)-1)
            dLevel = append(dLevel, levels[:i]...)
            dLevel = append(dLevel, levels[i+1:]...)
            if isSafe(dLevel) {
                safeFound = true
                break
            }
        }

        if safeFound {
            safe++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(safe)
}

func isSafe(levels []int) bool {
    var incr bool
    if levels[0] < levels[1] {
        incr = true
    } else {
        incr = false
    }

    for i:=0; i<len(levels)-1; i++ {
        firstNum := levels[i]
        secondNum := levels[i+1]

        if incr {
            if secondNum-firstNum <= 3 && secondNum-firstNum >= 1 {
                continue
            } else {
                return false
            }
        } else {
            if firstNum-secondNum <= 3 && firstNum-secondNum >=1 {
                continue
            } else {
                return false
            }
        }
    }

    return true
}
