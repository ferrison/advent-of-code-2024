package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Vec struct {
    x int
    y int
}

type Machine struct {
    a Vec
    b Vec
    prize Vec
}

var machines []Machine

func main() {
    file, err := os.Open("../input-13")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    machines = make([]Machine, 0)
    reBtn := regexp.MustCompile(`Button [AB]: X\+(\d*), Y\+(\d*)`)
    rePrize := regexp.MustCompile(`Prize: X=(\d*), Y=(\d*)`)
    for scanner.Scan() {
        btnALine := scanner.Text()
        scanner.Scan()
        btnBLine := scanner.Text()
        scanner.Scan()
        prizeLine := scanner.Text()
        scanner.Scan()

        btnAStrs := reBtn.FindStringSubmatch(btnALine)
        btnBStrs := reBtn.FindStringSubmatch(btnBLine)
        prizeStrs := rePrize.FindStringSubmatch(prizeLine)

        btnAX, _ := strconv.Atoi(btnAStrs[1])
        btnAY, _ := strconv.Atoi(btnAStrs[2])
        btnA := Vec{btnAX, btnAY}

        btnBX, _ := strconv.Atoi(btnBStrs[1])
        btnBY, _ := strconv.Atoi(btnBStrs[2])
        btnB := Vec{btnBX, btnBY}

        prizeX, _ := strconv.Atoi(prizeStrs[1])
        prizeY, _ := strconv.Atoi(prizeStrs[2])
        prize := Vec{prizeX+10000000000000, prizeY+10000000000000}

        machines = append(machines, Machine{btnA, btnB, prize})
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    sum := 0
    for _, machine := range machines {
        cost := process(machine)
        sum += cost
    }

    fmt.Println(sum)
}

func process(machine Machine) int {
    var a, b int
    commonDiv := machine.a.x*machine.b.y-machine.a.y*machine.b.x
    if commonDiv == 0 {
        return 0
    }

    aDen := machine.b.y*machine.prize.x-machine.b.x*machine.prize.y
    if aDen%commonDiv == 0 {
        a = aDen/commonDiv
    } else {
        return 0
    }

    bDen := machine.a.x*machine.prize.y - machine.a.y*machine.prize.x
    if bDen % commonDiv == 0 {
        b = bDen/commonDiv
    } else {
        return 0
    }

    if a < 0 || b < 0 {
        return 0
    }
    return a*3 + b
}
