package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

    sort.Ints(first)
    sort.Ints(second)

    var distances []int

    for i := 0; i < len(first); i++ {
        distance := int(math.Abs(float64(first[i] - second[i])))
        distances = append(distances, distance)
    }

    sum := 0

    for i := 0; i < len(distances); i++ {
        sum += distances[i]
    }

    fmt.Println("First: ")
    fmt.Println(first)
    fmt.Println("Second: ")
    fmt.Println(second)
    fmt.Println("Distances: ")
    fmt.Println(distances)
    fmt.Println("Sum: ")
    fmt.Println(sum)
}
