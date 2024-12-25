package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strconv"
)

type File struct {
    id int
    offset int
    size int
}

var disk []File
var checksum int

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
    offset := 0
    disk = make([]File, 0)
    for _, charB := range line {
        char := string(charB)
        size, _ := strconv.Atoi(char)

        if isFile {
            disk = append(disk, File{id: id, offset: offset, size: size})
            id++
        }

        offset += size
        isFile = !isFile
    }
    for i:=len(disk)-1; i>=0; {
        moved := moveFileToFreeSpace(&disk[i], i)
        if !moved {
            i--
        }
    }

    checksum = 0
    for _, file := range disk {
        for i:=file.offset; i<file.offset+file.size; i++ {
            checksum += i*file.id
        }
    }

    fmt.Println(checksum)
}

func moveFileToFreeSpace(fileToMove *File, fileIndex int) bool {
    for i:=0; i<fileIndex; i++ {
        file := disk[i]
        nextFile := disk[i+1]
        freeSpace := nextFile.offset - (file.offset + file.size)

        if freeSpace >= fileToMove.size {
            fileToMove.offset = file.offset + file.size
            newDisk := make([]File, 0)
            newDisk = append(newDisk, disk[:i+1]...)
            newDisk = append(newDisk, *fileToMove)
            newDisk = append(newDisk, disk[i+1:fileIndex]...)
            newDisk = append(newDisk, disk[fileIndex+1:]...)
            disk = newDisk
            return true
        }
    }
    return false
}
