package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Computer struct {
    name string
    connections map[*Computer]struct{}
}

var computers = map[string]*Computer{}

func main() {
    file, err := os.Open("../input-23")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        computerNames := strings.Split(line, "-")

        if _, ok := computers[computerNames[0]]; !ok {
            computers[computerNames[0]] = &Computer{computerNames[0], map[*Computer]struct{}{}}
        }

        if _, ok := computers[computerNames[1]]; !ok {
            computers[computerNames[1]] = &Computer{computerNames[1], map[*Computer]struct{}{}}
        }

        computer1 := computers[computerNames[0]]
        computer2 := computers[computerNames[1]]

        computer1.connections[computer2] = struct{}{}
        computer2.connections[computer1] = struct{}{}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    sets := getAllSets(computers)

    maxLen := 0
    var maxI int
    for i, set := range sets {
        if len(set) > maxLen {
            maxI = i
            maxLen = len(set)
        }
    }

    names := make([]string, 0)
    for _, c := range sets[maxI] {
        names = append(names, c.name)
    }
    sort.Strings(names)

    for i, name := range names {
        if i>0 {
            fmt.Print(",")
        }
        fmt.Print(name)
    }

    //printSets(sets)
}

func getAllSets(computers map[string]*Computer) [][]*Computer {
    sets := make([][]*Computer, 0)

    for _, computer := range computers {
        set := getLargestSet(computer)
        sets = append(sets, set)
    }
    return sets
}

func getLargestSet(computer *Computer) []*Computer {
    connections := make([]*Computer, 0)
    connections = append(connections, computer)
    for c := range computer.connections {
        connections = append(connections, c)
    }

    for n:=len(connections); n>0; n-- {
        combinations := getCombinations(connections, n)
        combs:
        for _, combination := range combinations {
            if !contains(combination, computer) {
                continue
            }

            for _, c := range combination {
                if !isConnected(c, combination) {
                    continue combs
                }
            }
            return combination
        }
    }

    return []*Computer{}
}

func contains(computers []*Computer, computer *Computer) bool {
    for _, c := range computers {
        if c == computer {
            return true
        }
    }
    return false
}

func isConnected(c *Computer, connections []*Computer) bool {
    for _, conn := range connections {
        if conn==c {
            continue
        }
        if _, ok := c.connections[conn]; !ok {
            return false
        }
    }
    return true
}

func getCombinations(computers []*Computer, n int) [][]*Computer {
    if n == 1 {
        res := [][]*Computer{}
        for _, c := range computers {
            res = append(res, []*Computer{c})
        }
        return res
    }

    res := make([][]*Computer, 0)
    for i:=0; i<len(computers); i++ {
        remainingComputers := computers[i+1:]
        for _, combination := range getCombinations(remainingComputers, n-1) {
            newCombination := append([]*Computer{computers[i]}, combination...)
            res = append(res, newCombination)
        }
    }
    return res
}

func printSets(sets [][]*Computer) {
    for _, set := range sets {
        for _, c := range set {
            fmt.Printf("%s, ", c.name)
        }
        fmt.Println()
    }
}
