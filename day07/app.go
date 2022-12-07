package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var lineLength int64

func getInputAsArray(filePath string) []string {
	dat, err := os.ReadFile(filePath)
	check(err)
	lines := strings.Split(string(dat), "\n")
	lineLength = int64(len(lines))
	return lines
}

func scanDir(lines []string, lineIndex *int64, solution *int64) int64 {

	var size int64
	line := lines[*lineIndex]
	var dirName string

	if strings.HasPrefix(line, "$ cd") && !strings.HasPrefix(line, "$ cd ..") {
		dirName = line[5:]
	}

	if line[0] >= '0' && line[0] <= '9' {
		fileSize, _ := strconv.ParseInt(strings.Split(line, " ")[0], 10, 64)
		size += fileSize
	}

	*lineIndex++
	if strings.HasPrefix(line, "$ cd ..") {
		return size
	}

	if *lineIndex < lineLength {
		size += scanDir(lines, lineIndex, solution)
		// if dirName != "" && size <= 100000 {
		if dirName != "" {
			if size < 100000 {
				// println(dirName, size)
				*solution += size
			}

			if *lineIndex < lineLength {
				line = lines[*lineIndex]
				if !strings.HasPrefix(line, "$ cd ..") {
					size += scanDir(lines, lineIndex, solution)
				} else {
					*lineIndex++
				}
			}
		}
	}
	return size
}

func part1(lines []string) {

	var lineIndex int64 = 0
	var solution int64 = 0
	var rootSize int64 = 0
	for lineIndex < lineLength {
		rootSize += scanDir(lines, &lineIndex, &solution)
	}
	if rootSize <= 100000 {
		solution += rootSize
	}
	println(solution)
}

func part2(lines []string) {
	fmt.Println("TODO")
}

func main() {

	lines := getInputAsArray("input.txt")

	part := os.Getenv("part")
	if part == "part2" {
		part2(lines)
	} else {
		part1(lines)
	}
}
