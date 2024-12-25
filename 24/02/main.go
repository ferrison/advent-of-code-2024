package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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
var swaps = []string{}

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
        
    for i:=2; ;i++ {
        out := fmt.Sprintf("z%02d", i)
        inx := fmt.Sprintf("x%02d", i)
        iny := fmt.Sprintf("y%02d", i)

        _, okz := outputs[out]
        _, okx := vals[inx]
        _, oky := vals[iny]

        if okz && okx && oky {
            isCorrect(out)
        } else {
            break
        }
    }

    sort.Strings(swaps)
    for i, s := range swaps {
        if i > 0 {
            fmt.Print(",")
        }
        fmt.Print(s)
    }
    fmt.Println()
}

func isCorrect(out string) {
    node := outputs[out]
    outVal, _ := strconv.Atoi(out[1:])

    if node.op != XOR {
        wantedXorOperand := findWanted(XOR, fmt.Sprintf("x%02d", outVal), fmt.Sprintf("y%02d", outVal))
        wanted := findWanted(XOR, wantedXorOperand, "")
        swap(out, wanted)
        node = outputs[out]
    }

    isIn1XOR := isOp(XOR, node.in1, outVal)
    isIn2XOR := isOp(XOR, node.in2, outVal)
    if !isIn1XOR && !isIn2XOR {
        wantedXorOperand := findWanted(XOR, fmt.Sprintf("x%02d", outVal), fmt.Sprintf("y%02d", outVal))
        wanted := findWanted(XOR, wantedXorOperand, "")
        if wanted != "" {
            swap(out, wanted)
        } else {
            isIn1ProbablyCarry := probablyIsCarry(node.in1, outVal)

            if isIn1ProbablyCarry {
                swap(node.in2, wantedXorOperand)
                isIn2XOR = true
            } else {
                swap(node.in1, wantedXorOperand)
                isIn1XOR = true
            }
        }
        node = outputs[out]
    }

    if isIn1XOR {
        isCarryNode(node.in2, outVal)
    } else {
        isCarryNode(node.in1, outVal)
    }
}

func swap(out1 string, out2 string) {
    node1 := outputs[out1]
    node2 := outputs[out2]

    outputs[out1] = node2
    outputs[out2] = node1

    fmt.Printf("%s swapped with %s\n", out1, out2)
    swaps = append(swaps, out1, out2)
}

func isOp(op string, out string, val int) bool {
    n := outputs[out]
    valStr := fmt.Sprintf("%02d", val)
    in1 := "x" + valStr
    in2 := "y" + valStr

    return n.op == op && (n.in1 == in1 && n.in2 == in2 || n.in1 == in2 && n.in2 == in1)
}

func probablyIsCarry(out string, carryFor int) bool {
    if carryFor == 1 {
        if !isOp(AND, out, 0) {
            return false
        }
        return true
    }

    node := outputs[out]

    if node.op != OR {
        return false
    }

    isIn1AND := isOp(AND, node.in1, carryFor-1)
    isIn2AND := isOp(AND, node.in2, carryFor-1)

    if !isIn1AND && !isIn2AND {
        return false
    }
    
    return true
}

func isCarryNode(out string, carryFor int) {
    if carryFor == 1 {
        if !isOp(AND, out, 0) {
            fmt.Printf("%s: operands must be AND'ed\n", out)
        }
        return
    }

    node := outputs[out]

    if node.op != OR {
        wantedAndOperand := findWanted(AND, fmt.Sprintf("x%02d", carryFor-1), fmt.Sprintf("y%02d", carryFor-1))
        wanted := findWanted(OR, wantedAndOperand, "")
        swap(out, wanted)
        node = outputs[out]
    }

    isIn1AND := isOp(AND, node.in1, carryFor-1)
    isIn2AND := isOp(AND, node.in2, carryFor-1)

    if !isIn1AND && !isIn2AND {
        wantedAndOperand := findWanted(AND, fmt.Sprintf("x%02d", carryFor-1), fmt.Sprintf("y%02d", carryFor-1))
        wanted := findWanted(OR, wantedAndOperand, "")
        if wanted != "" {
            swap(out, wanted)
        } else {
            isIn1ProbablySubcarry := probablyIsSubcarry(node.in1, carryFor)

            if isIn1ProbablySubcarry {
                swap(node.in2, wantedAndOperand)
                isIn2AND = true
            } else {
                swap(node.in1, wantedAndOperand)
                isIn1AND = true
            }
        }
        node = outputs[out]
    }
    
    if isIn1AND {
        isSubcarryNode(node.in2, carryFor)
    } else {
        isSubcarryNode(node.in1, carryFor)
    }
}

func probablyIsSubcarry(out string, carryFor int) bool {
    node := outputs[out]

    if node.op != AND {
        return false
    }

    isIn1XOR := isOp(XOR, node.in1, carryFor-1)
    isIn2XOR := isOp(XOR, node.in2, carryFor-1)
    if !isIn1XOR && !isIn2XOR {
        return false
    }

    return true
}

func isSubcarryNode(out string, carryFor int) {
    node := outputs[out]

    if node.op != AND {
        wantedXorOperand := findWanted(XOR, fmt.Sprintf("x%02d", carryFor-1), fmt.Sprintf("y%02d", carryFor-1))
        wanted := findWanted(AND, wantedXorOperand, "")
        swap(out, wanted)
        node = outputs[out]
    }

    isIn1XOR := isOp(XOR, node.in1, carryFor-1)
    isIn2XOR := isOp(XOR, node.in2, carryFor-1)
    if !isIn1XOR && !isIn2XOR {
        wantedXorOperand := findWanted(XOR, fmt.Sprintf("x%02d", carryFor-1), fmt.Sprintf("y%02d", carryFor-1))
        wanted := findWanted(OR, wantedXorOperand, "")
        if wanted != "" {
            swap(out, wanted)
        } else {
            isIn1ProbablyCarry := probablyIsCarry(node.in1, carryFor-1)

            if isIn1ProbablyCarry {
                swap(node.in2, wantedXorOperand)
                isIn2XOR = true
            } else {
                swap(node.in1, wantedXorOperand)
                isIn1XOR = true
            }
        }
        node = outputs[out]
    }

    if isIn1XOR {
        isCarryNode(node.in2, carryFor-1)
    } else {
        isCarryNode(node.in1, carryFor-1)
    }
}

func findWanted(op string, in1 string, in2 string) string {
    for _, n := range outputs {
        if in2 == "" {
            if n.op == op && (n.in1 == in1 || n.in2 == in1) {
                return n.out
            }
        }

        if n.op == op && (n.in1 == in1 && n.in2 == in2 || n.in1 == in2 && n.in2 == in1) {
            return n.out
        }
    }
    return ""
}
