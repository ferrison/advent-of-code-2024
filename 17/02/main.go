package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

var program []int

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
    scanner.Scan()
    scanner.Scan()
    scanner.Scan()

    programLine := scanner.Text()

    registerAStr := reRegister.FindStringSubmatch(registerALine)
    instrStrs := reProgram.FindAllString(programLine, -1)

    registerA, _ := strconv.Atoi(registerAStr[1])

    program = make([]int, 0)
    for _, instrStr := range instrStrs {
        instr, _ := strconv.Atoi(instrStr)
        program = append(program, instr)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(registerA)

    //35184372088832 35184372088832*8-1
    //175921862088832 -> output[15] == 0
    //start := 35184372088832
    fmt.Println(process(15, int(math.Pow(8, 15))))
}

func process(i int, curr int) bool {
    if i<0 {
        return true
    }
    for j:=0; ; j++ {
        fmt.Printf("i: %d\n", i)
        fmt.Printf("j: %d\n", j)
        initA := curr + int(math.Pow(8, float64(i)))*j
        A := initA
        fmt.Printf("A: %d\n", A)
        output := make([]int, 0)
        for A > 0 {
            //output = append(output, (((A%8)^5)^(A/(int(math.Pow(2, float64((A%8)^1))))))%8)
            //output = append(output, ((A%8^5)^(A>>(A%8^1)))%8)
            output = append(output, (A%8^5)^((A>>(A%8^1))%8))
            A /= 8
        }
        fmt.Println(program)
        fmt.Println(output)
        fmt.Println()
        if !reflect.DeepEqual(output[i+1:], program[i+1:]) {
            return false
        }
        if reflect.DeepEqual(output[i:], program[i:]) {
            res := process(i-1, initA)
            if res {
                return true
            }
        }
    }
}
