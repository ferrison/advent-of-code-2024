package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var sum int

func main() {
    file, err := os.Open("../input-07")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    sum = 0
    for scanner.Scan() {
        line := scanner.Text()

        resAndValsStr := strings.Split(line, ":")

        resStr := resAndValsStr[0]
        valsStr:= resAndValsStr[1]

        res, _ := strconv.Atoi(resStr)

        var vals []int
        for _, valStr := range strings.Fields(valsStr) {
            val, _ := strconv.Atoi(valStr)
            vals = append(vals, val)
        }

        check(vals[0], res, vals[1:])
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(sum)
}

func check(curr int, res int, vals []int) bool {
    if len(vals) == 0 {
        if res == curr {
            sum += res
            return true
        }
        return false
    }
    if check(curr+vals[0], res, vals[1:]) {
        return true
    }
    return check(curr*vals[0], res, vals[1:])
}
