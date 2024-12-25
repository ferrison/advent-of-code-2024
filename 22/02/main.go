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

    for scanner.Scan() {
        line := scanner.Text()

        num, _ := strconv.Atoi(line)
        calcNth(num)
    }

    maxV := 0
    var maxK Seq
    for k, v := range seqToPrices {
        if v > maxV {
            maxK = k
            maxV = v
        }
    }
    fmt.Println(maxK)
    fmt.Println(seqToPrices[maxK])
}

type Seq struct {
    a int
    b int
    c int
    d int
}

var seqToPrices = map[Seq]int{}

func calcNth(num int) {
    prices := make([]int, 0)
    prev := num % 10

    localSeqToPrice := make(map[Seq]int)

    for i:=0; i<2000+1; i++ {
        num = mix(num, num * 64)
        num = prune(num)

        num = mix(num, num/32)
        num = prune(num)

        num = mix(num, num * 2048)
        num = prune(num)

        price := num % 10
        diff := price - prev
        prices = append(prices, diff)

        if len(prices) >= 4 {
            fourLastPrices := prices[len(prices)-4:]
            seq := Seq{fourLastPrices[0], fourLastPrices[1], fourLastPrices[2], fourLastPrices[3]}
            if _, ok := localSeqToPrice[seq]; !ok {
                localSeqToPrice[seq] = price
            }
        }

        prev = price
    }

    for k, v := range localSeqToPrice {
        seqToPrices[k] += v
    }
}

func mix(num int, op int) int {
    return num ^ op
}

func prune(num int) int {
    return num % 16777216
}
