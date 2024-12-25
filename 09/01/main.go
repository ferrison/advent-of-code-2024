package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strconv"
)

var disk []int

func main() {
    file, err := os.Open("../input-09")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var line string
    for scanner.Scan() {
        line = scanner.Text()
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    isFile := true
    id := 0
    for _, charB := range line {
        char := string(charB)
        size, _ := strconv.Atoi(char)
        for i:=0; i<size; i++ {
            var aId int
            if isFile {
                aId = id
            } else {
                aId = -1
            }
            disk = append(disk, aId)
        }
        if isFile {
            id++
        }
        isFile = !isFile
    }

    i := 0
    j := len(disk)-1
    checksum := 0
    for {
        if i-1==j {
            break
        }
        if disk[i] != -1 {
            fmt.Printf("%d:%d\n", i, disk[i])
            checksum += i*disk[i]
        } else {
            for ;disk[j]==-1;j-- {}
            fmt.Printf("%d:%d\n", i, disk[j])
            checksum += i*disk[j]
            j--
        }
        i++
    }

    fmt.Println(checksum)
}
