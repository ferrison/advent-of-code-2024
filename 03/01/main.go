package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
    data, err := os.ReadFile("../input-03")
    if err != nil {
        log.Fatal(err)
    }
    content := string(data)
    re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

    matches := re.FindAllStringSubmatch(content, -1)

    const (
        mul = iota
        do = iota
        dont = iota
    )

    calc := true
    sum := 0
    for _, match := range matches {
        full := match[0]
        
        var instr int
        if strings.HasPrefix(full, "mul") {
            instr = mul
        }
        if strings.HasPrefix(full, "do(") {
            instr = do
        }
        if strings.HasPrefix(full, "don") {
            instr = dont
        }

        switch instr {
        case do:
            calc = true
            continue
        case dont:
            calc = false
            continue
        case mul:
            if calc {
                a, err1 := strconv.Atoi(match[1])
                b, err2 := strconv.Atoi(match[2])
                if err1 != nil || err2 != nil {
                    log.Fatal(err1, err2)
                }
                sum += a*b
            }
        }
    }


    fmt.Println(sum)
}
