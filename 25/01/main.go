package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var locks = [][5]int{}
var keys = [][5]int{}

func main() {
    file, err := os.Open("../input-25")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    grid := [7][5]string{}
    for scanner.Scan() {
        for row:=0; row<7; row++ {
            line := scanner.Text()
            for col, elR := range line {
                el := string(elR)
                grid[row][col] = el
            }
            scanner.Scan()
        }

        vals := [5]int{}
        for col:=0; col<5; col++ {
            count := 0
            for row:=0; row<7; row++ {
                if grid[row][col] == "#" {
                    count++
                }
            }
            vals[col] = count-1
        }
        if grid[0][0] == "#" {
            locks = append(locks, vals)
        } else {
            keys = append(keys, vals)
        }

    }

    count := 0
    for _, lock := range locks {
        for _, key := range keys {
            if fit(lock, key) {
                count++
            }
        }
    }

    fmt.Println(count)
}

func fit(lock, key [5]int) bool {
    for i:=0; i<5; i++ {
        if lock[i]+key[i] > 5 {
            return false
        }
    }
    return true
}
