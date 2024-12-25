package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
    x int
    y int
}

func (p Point) Add(other Point) Point {
    return Point{
        x: p.x + other.x,
        y: p.y + other.y,
    }
}

func (p Point) Sub(other Point) Point {
    return Point{
        x: p.x - other.x,
        y: p.y - other.y,
    }
}

func (p Point) Mul2() Point {
    return Point{
        x: p.x*2,
        y: p.y*2,
    }
}

var antennasMap map[string][]Point
var antinodes map[Point]struct{}
var w, h int

func main() {
    file, err := os.Open("../input-08")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    y := 0
    antennasMap := make(map[string][]Point)
    for scanner.Scan() {
        line := scanner.Text()
        for x, charR := range line {
            char := string(charR)
            if char != "." {
                if _, ok := antennasMap[char]; !ok {
                    antennasMap[char] = make([]Point, 0)
                }
                antennasMap[char] = append(antennasMap[char], Point{x, y})
            }
        }

        y++
        w = len(line)
        h = y
    }

    antinodes = make(map[Point]struct{})
    for _, antennas := range antennasMap {
        for i:=0; i<len(antennas); i++ {
            antenna := antennas[i]
            otherAntennas := make([]Point, 0, len(antennas)-1)
            otherAntennas = append(otherAntennas, antennas[:i]...)
            otherAntennas = append(otherAntennas, antennas[i+1:]...)
            calcAntinodes(antenna, otherAntennas)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(len(antinodes))
}

func calcAntinodes(antenna Point, otherAntennas []Point) {
    for _, otherAntenna := range otherAntennas {
        vec := otherAntenna.Sub(antenna)
        antinode := antenna.Add(vec.Mul2())
        if antinode.x >= 0 && antinode.x < w && antinode.y >= 0 && antinode.y < h {
            antinodes[antinode] = struct{}{}
        }
    }
}
