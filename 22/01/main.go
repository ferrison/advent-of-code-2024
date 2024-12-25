package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
    file, err := os.Open("../input-22")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    sum := 0
    for scanner.Scan() {
        line := scanner.Text()

        num, _ := strconv.Atoi(line)
        sum += calcNth(num)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(sum)
}

func calcNth(num int) int {
    for i:=0; i<2000; i++ {
        num = mix(num, num * 64)
        num = prune(num)

        num = mix(num, num/32)
        num = prune(num)

        num = mix(num, num * 2048)
        num = prune(num)
    }
    return num
}

func mix(num int, op int) int {
    return num ^ op
}

func prune(num int) int {
    return num % 16777216
}
