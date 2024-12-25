package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var beforeMap map[int]map[int]struct{}

func main() {
    file, err := os.Open("../input-05")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    beforeMap = make(map[int]map[int]struct{})

    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            break
        }

        orderStr:= strings.Split(line, "|")

        before, _ := strconv.Atoi(orderStr[0])
        after, _ := strconv.Atoi(orderStr[1])

        if _, ok := beforeMap[after]; !ok {
            beforeMap[after] = make(map[int]struct{})
        }

        beforeSet := beforeMap[after]
        beforeSet[before] = struct{}{}
    }

    for k, v := range beforeMap {
        fmt.Printf("%d:", k)
        var nums []int

        for num := range v {
            nums = append(nums, num)
        }
        fmt.Printf("%d\n", nums)
    }

    sum := 0
    for scanner.Scan() {
        line := scanner.Text()

        numsStr := strings.Split(line, ",")

        var nums []int

        for _, numStr := range numsStr {
            num, _ := strconv.Atoi(numStr)
            nums = append(nums, num)
        }
        
        if !check(nums) {
            correctNums := correct(nums)
            sum += correctNums[len(nums)/2]
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(sum)
}

func check(nums []int) bool {
    var seenNums []int
    for _, num := range nums {
        for _, seenNum := range seenNums {
            if beforeSet, ok := beforeMap[seenNum]; ok {
                if _, ok := beforeSet[num]; ok {
                    return false
                }
            }
        }
        seenNums = append(seenNums, num)
    }
    return true
}

func correct(nums []int) []int {
    for i:=0; i<len(nums); i++ {
        num := nums[i]
        for j:=0; j<i; j++ {
            seenNum := nums[j]
            beforeSet := beforeMap[seenNum]
            if _, ok := beforeSet[num]; ok {
                newNums := make([]int, 0, len(nums))
                newNums = append(newNums, nums[:j]...)
                newNums = append(newNums, num)
                newNums = append(newNums, nums[j:i]...)
                newNums = append(newNums, nums[i+1:]...)
                nums = newNums
                break
            }
        }
    }
    return nums
}
