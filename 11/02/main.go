package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Args struct {
    val int
    lvl int
}

var stones []int
var cache map[Args]int
const MAX_LVL=75

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

    cache = make(map[Args]int)
    stonesLen := 0
    for _, val := range stones {
        stonesLen += count(val, 0)
    }

    fmt.Println(stonesLen)
}

func count(val int, lvl int) int {
    if lvl == MAX_LVL {
        return 1
    }
    if c, ok := cache[Args{val, lvl}]; ok {
        return c
    }
    if val == 0 {
        c := count(1, lvl+1)
        cache[Args{val, lvl}] = c
        return c
    }
    valStr := strconv.Itoa(val)
    if len(valStr) % 2 == 0 {
        val1, _ := strconv.Atoi(valStr[:len(valStr)/2])
        val2, _ := strconv.Atoi(valStr[len(valStr)/2:])
        c := count(val1, lvl+1) + count(val2, lvl+1)
        cache[Args{val, lvl}] = c
        return c
    }
    c := count(val * 2024, lvl+1)
    cache[Args{val, lvl}] = c
    return c
}
