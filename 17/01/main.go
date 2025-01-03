package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
    file, err := os.Open("../input-17")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    reRegister := regexp.MustCompile(`Register [ABC]: (\d*)`)
    reProgram := regexp.MustCompile(`\d+`)

    scanner.Scan()
    registerALine := scanner.Text()
    scanner.Scan()
    registerBLine := scanner.Text()
    scanner.Scan()
    registerCLine := scanner.Text()
    scanner.Scan()
    scanner.Scan()

    programLine := scanner.Text()

    registerAStr := reRegister.FindStringSubmatch(registerALine)
    registerBStr := reRegister.FindStringSubmatch(registerBLine)
    registerCStr := reRegister.FindStringSubmatch(registerCLine)
    instrStrs := reProgram.FindAllString(programLine, -1)

    registerA, _ := strconv.Atoi(registerAStr[1])
    registerB, _ := strconv.Atoi(registerBStr[1])
    registerC, _ := strconv.Atoi(registerCStr[1])

    program := make([]int, 0)
    for _, instrStr := range instrStrs {
        instr, _ := strconv.Atoi(instrStr)
        program = append(program, instr)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    computer := Computer{
        A: registerA,
        B: registerB,
        C: registerC,
        program: program,
        instrP: 0,
        output: make([]int, 0),
    }

    for computer.tick() {}
    
    for i:=0; i<=len(computer.output)-1; i++ {
        fmt.Print(computer.output[i])
        if i != len(computer.output)-1 {
            fmt.Print(",")
        }
    }
}

type Computer struct {
    A int
    B int
    C int
    program []int
    instrP int
    output []int
}

func (c *Computer) combo(op int) int {
    switch {
    case 0 <= op && op <= 3:
        return op
    case op == 4:
        return c.A
    case op == 5:
        return c.B
    case op == 6:
        return c.C
    case op == 7:
        log.Fatal("Unexpected combo operand")
    }
    log.Fatal("Unexpected combo function branch")
    return 0
}

func (c *Computer) tick() bool {
    if c.instrP >= len(c.program) {
        return false
    }
    opCode := c.program[c.instrP]
    op := c.program[c.instrP+1]
    fmt.Printf("opCode: %d\n", opCode)
    fmt.Printf("op: %d\n", op)
    fmt.Printf("A: %d\n", c.A)
    fmt.Printf("B: %d\n", c.B)
    fmt.Printf("C: %d\n", c.C)
    fmt.Printf("instrP: %d\n", c.instrP)
    fmt.Println()

    switch opCode {
    case 0:
        c.adv(op)
    case 1:
        c.bxl(op)
    case 2:
        c.bst(op)
    case 3:
        c.jnz(op)
    case 4:
        c.bxc(op)
    case 5:
        c.out(op)
    case 6:
        c.bdv(op)
    case 7:
        c.cdv(op)
    }

    return true
}

func (c *Computer) adv(op int) {
    res := c.A/int(math.Pow(2, float64(c.combo(op))))
    c.A = res
    c.instrP += 2
}

func (c *Computer) bxl(op int) {
    res := c.B ^ int(op)
    c.B = res
    c.instrP += 2
}

func (c *Computer) bst(op int) {
    res := c.combo(op)%8
    c.B = res
    c.instrP += 2
}

func (c *Computer) jnz(op int) {
    if c.A == 0 {
        c.instrP += 2
        return
    }
    c.instrP = int(op)
}

func (c *Computer) bxc(_ int) {
    res := c.B ^ c.C
    c.B = res
    c.instrP += 2
}

func (c *Computer) out(op int) {
    res := c.combo(op) % 8
    c.output = append(c.output, res)
    c.instrP += 2
}

func (c *Computer) bdv(op int) {
    res := c.A/int(math.Pow(2, float64(c.combo(op))))
    c.B = res
    c.instrP += 2
}

func (c *Computer) cdv(op int) {
    res := c.A/int(math.Pow(2, float64(c.combo(op))))
    c.C = res
    c.instrP += 2
}
