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
    var first []int;
    var second []int;
    file, err := os.Open("input-01")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        numStrs := strings.Split(line, "   ")

        num1, err := strconv.Atoi(numStrs[0])
        if err != nil {
            log.Fatal(err)
        }

        num2, err := strconv.Atoi(numStrs[1])
        if err != nil {
            log.Fatal(err)
        }

        first = append(first, num1)
        second = append(second, num2)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    countSecond := make(map[int]int)
    for i := 0; i < len(second); i++ {
        id := second[i]
        if c, ok := countSecond[id]; ok {
            countSecond[id] = c+1
        } else {
            countSecond[id] = 1
        }
    }

    similarityScore := 0

    for i:=0; i<len(first); i++ {
        id := first[i]
        if c, ok := countSecond[id]; ok {
            similarityScore += id*c
        }
    }

    fmt.Println("First: ")
    fmt.Println(first)
    fmt.Println("Second: ")
    fmt.Println(second)
    fmt.Println("countSecond: ")
    fmt.Println(countSecond)
    fmt.Println("similarityScore: ")
    fmt.Println(similarityScore)
}
