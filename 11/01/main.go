package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var stones []int

func main() {
    lineB, err := os.ReadFile("../input-11")
    line := string(lineB)

    if err != nil {
        log.Fatal(err)
    }

    stones = make([]int, 0)
    for _, valStr := range strings.Fields(line) {
        val, _ := strconv.Atoi(valStr)
        stones = append(stones, val)
    }

    for i:=0; i<25; i++ {
        blink()
    }

    fmt.Println(len(stones))
}

func blink() {
    lenStones := len(stones)
    for i:=0; i<lenStones; i++ {
        val := stones[i]
        if val == 0 {
            stones[i] = 1
            continue
        }
        valStr := strconv.Itoa(val)
        if len(valStr) % 2 == 0 {
            val1, _ := strconv.Atoi(valStr[:len(valStr)/2])
            val2, _ := strconv.Atoi(valStr[len(valStr)/2:])

            stones = append(stones[:i], append([]int{val1, val2}, stones[i+1:]...)...)
            lenStones++
            i++
            continue
        }
        stones[i] = val * 2024
    }
}
