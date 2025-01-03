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

func (v Vec) Add(other Vec) Vec {
    return Vec{
        v.x+other.x,
        v.y+other.y,
    }
}

func (v Vec) Mul(n int) Vec {
    return Vec{
        v.x*n,
        v.y*n,
    }
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
        prize := Vec{prizeX, prizeY}

        machines = append(machines, Machine{btnA, btnB, prize})
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    sum := 0
    for _, machine := range machines {
        minCost = 1000
        cache = make(map[Args]Res)
        cost, win := process(machine, 0, 0)
        if win {
            sum += cost
        }
    }

    fmt.Println(sum)
}

var minCost int

type Args struct {
    a int
    b int
}

type Res struct {
    cost int
    win bool
}

var cache map[Args]Res

func process(machine Machine, aPressed int, bPressed int) (int, bool) {
    currCost := aPressed*3 + bPressed
    if currCost > minCost {
        cache[Args{aPressed, bPressed}] = Res{currCost, false}
        return currCost, false
    }
    if aPressed > 100 || bPressed > 100 {
        cache[Args{aPressed, bPressed}] = Res{currCost, false}
        return currCost, false
    }
    currPos := machine.a.Mul(aPressed).Add(machine.b.Mul(bPressed))
    if currPos.x > machine.prize.x || currPos.y > machine.prize.y {
        cache[Args{aPressed, bPressed}] = Res{currCost, false}
        return currCost, false
    }
    if currPos == machine.prize {
        minCost = min(minCost, currCost)
        cache[Args{aPressed, bPressed}] = Res{currCost, true}
        return currCost, true
    }
    if res, ok := cache[Args{aPressed, bPressed}]; ok {
        return res.cost, res.win
    }
    pressB, winB := process(machine, aPressed, bPressed+1)
    pressA, winA := process(machine, aPressed+1, bPressed)
    if winA && winB {
        cost := min(pressA, pressB)
        cache[Args{aPressed, bPressed}] = Res{cost, false}
        return cost, true
    }
    if winA {
        cache[Args{aPressed, bPressed}] = Res{pressA, false}
        return pressA, true
    }
    if winB {
        cache[Args{aPressed, bPressed}] = Res{pressB, false}
        return pressB, true
    }
    cache[Args{aPressed, bPressed}] = Res{currCost, false}
    return currCost, false
}
