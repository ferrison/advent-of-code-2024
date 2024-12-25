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

    linesCount := 0
    unsafe := 0
    for scanner.Scan() {
        linesCount++
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
                    unsafe++
                    break
                }
            } else {
                if firstNum-secondNum <= 3 && firstNum-secondNum >=1 {
                    continue
                } else {
                    unsafe++
                    break
                }
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(linesCount-unsafe)
}
