package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var towels []string

func main() {
    file, err := os.Open("../input-19")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    scanner.Scan()
    towelsLine := scanner.Text()

    scanner.Scan()

    towels = make([]string, 0)
    for _, towel := range strings.Split(towelsLine, ", ") {
        towels = append(towels, towel)
    }

    patterns := make([]string, 0)
    for scanner.Scan() {
        line := scanner.Text()
        patterns = append(patterns, line)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    sum := 0
    cache = make(map[string]bool)
    for i, pattern := range patterns {
        res := isPossible(pattern)
        if res {
            sum++
        }
        fmt.Println(i)
    }

    fmt.Println(sum)
}

var cache map[string]bool

func isPossible(pattern string) bool {
    if pattern == "" {
        return true
    }

    if res, ok := cache[pattern]; ok {
        return res
    }

    for _, towel := range towels {
        if strings.HasPrefix(pattern, towel) {
            newPattern, _ := strings.CutPrefix(pattern, towel)
            if isPossible(newPattern) {
                cache[pattern] = true
                return true
            }
        }
    }

    cache[pattern] = false
    return false
}
