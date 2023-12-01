package main

import (
	"os"
	"sort"
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

func scanDir(lines []string, lineIndex *int64, solutionPart1 *int64, solutionPart2 *[]int64) int64 {

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
		size += scanDir(lines, lineIndex, solutionPart1, solutionPart2)
		if dirName != "" {
			// println(dirName, size)
			if size < 100000 {
				// println(dirName, size)
				*solutionPart1 += size
			}
			*solutionPart2 = append(*solutionPart2, size)

			if *lineIndex < lineLength {
				line = lines[*lineIndex]
				if !strings.HasPrefix(line, "$ cd ..") {
					size += scanDir(lines, lineIndex, solutionPart1, solutionPart2)
				} else {
					*lineIndex++
				}
			}
		}
	}
	return size
}

func solve(part string, lines []string) {

	var lineIndex int64 = 0
	var solutionPart1 int64 = 0
	var solutionPart2 []int64
	var rootSize int64 = 0
	rootSize += scanDir(lines, &lineIndex, &solutionPart1, &solutionPart2)

	// part 2
	if part == "part2" {
		var totalDiskSpace int64 = 70000000
		var requiredSpace int64 = 30000000
		deleteLimit := requiredSpace - (totalDiskSpace - rootSize)
		sort.Slice(solutionPart2, func(i, j int) bool { return solutionPart2[i] < solutionPart2[j] })
		for i, size := range solutionPart2 {
			if size > deleteLimit {
				println(solutionPart2[i])
				break
			}
		}
	} else {
		println(solutionPart1)
	}
}

func main() {

	lines := getInputAsArray("input.txt")

	part := os.Getenv("part")
	solve(part, lines)
}
