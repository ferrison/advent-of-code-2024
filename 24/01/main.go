package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
    AND = "AND"
    OR = "OR"
    XOR = "XOR"
)

type Node struct {
    op string
    in1 string
    in2 string
    out string
}

var outputs = map[string]Node{}
var vals = map[string]bool{}

func main() {
    file, err := os.Open("../input-24")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    reInitVals := regexp.MustCompile(`([xy]\d\d): (\d)`)
    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            break
        }

        matches := reInitVals.FindStringSubmatch(line)
        name := matches[1]
        valStr := matches[2]
        var val bool
        if valStr == "1" {
            val = true
        } else {
            val = false
        }

        vals[name] = val
    }

    reGates := regexp.MustCompile(`(\w{3}) (AND|OR|XOR) (\w{3}) -> (\w{3})`)
    for scanner.Scan() {
        line := scanner.Text()

        matches := reGates.FindStringSubmatch(line)
        in1 := matches[1]
        op := matches[2]
        in2 := matches[3]
        out := matches[4]

        outputs[out] = Node{op, in1, in2, out}
    }
        
    res := ""
    for i:=0; ;i++ {
        out := fmt.Sprintf("z%02d", i)
        if _, ok := outputs[out]; ok {
            outVal := calc(out)
            if outVal {
                res = "1" + res
            } else {
                res = "0" + res
            }
        } else {
            break
        }
    }

    fmt.Println(res)
    res10, _ := strconv.ParseUint(res, 2, 64)
    fmt.Println(res10)
}

func calc(out string) bool {
    if outVal, ok := vals[out]; ok {
        return outVal
    }

    node := outputs[out]
    var res bool
    switch node.op {
    case AND:
        res = calc(node.in1) && calc(node.in2)
    case OR:
        res = calc(node.in1) || calc(node.in2)
    case XOR:
        res = calc(node.in1) != calc(node.in2)
    }

    vals[out] = res
    return res
}
