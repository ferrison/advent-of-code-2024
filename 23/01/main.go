package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
    sum := 0
    for _, set := range sets {
        for c := range set {
            if strings.HasPrefix(c.name, "t") {
                sum++
                break
            }
        }
    }
    fmt.Println(sum)
}

func getAllSets(computers map[string]*Computer) []map[*Computer]struct{} {
    processed := make(map[*Computer]struct{})
    sets := make([]map[*Computer]struct{}, 0)

    for _, computer := range computers {
        computerSets := get3Sets(computer, processed)
        sets = append(sets, computerSets...)
        processed[computer] = struct{}{}
    }
    return sets
}

func get3Sets(computer *Computer, processed map[*Computer]struct{}) []map[*Computer]struct{} {
    sets := make([]map[*Computer]struct{}, 0)

    connections := make([]*Computer, 0, len(computer.connections))
    for c := range computer.connections {
        connections = append(connections, c)
    }

    for i:=0; i<len(connections); i++ {
        for j:=i+1; j<len(connections); j++ {
            c1 := connections[i]
            c2 := connections[j]

            if _, ok := processed[c1]; ok {
                continue
            }

            if _, ok := processed[c2]; ok {
                continue
            }

            if _, ok := c1.connections[c2]; ok {
                sets = append(sets, map[*Computer]struct{}{
                    computer: {},
                    c1: {},
                    c2: {},
                })
            }
        }
    }

    return sets
}

func printSets(sets []map[*Computer]struct{}) {
    for _, set := range sets {
        for c := range set {
            fmt.Printf("%s, ", c.name)
        }
        fmt.Println()
    }
}
