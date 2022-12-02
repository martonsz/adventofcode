package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getInputAsArray(filePath string) []string {
	dat, err := os.ReadFile(filePath)
	check(err)
	return strings.Split(string(dat), "\n")
}

func printCWD() {
	os.Chdir("/Users")
	newDir, _ := os.Getwd()
	fmt.Printf("Current Working Direcotry: %s\n", newDir)
}

func part1(lines []string) {

	totalScore := 0
	for _, line := range lines {

		elf := moveMap[line[0]]
		me := moveMap[line[2]]

		totalScore += me + 1

		if (elf+1)%3 == me {
			totalScore += 6
		} else if elf == me {
			totalScore += 3
		}

	}
	fmt.Println(totalScore)
}

func part2(lines []string) {
	fmt.Println("TODO")
}

var moveMap = make(map[byte]int)

func main() {
	moveMap['A'] = 0
	moveMap['B'] = 1
	moveMap['C'] = 2
	moveMap['X'] = 0
	moveMap['Y'] = 1
	moveMap['Z'] = 2

	lines := getInputAsArray("input.txt")

	part := os.Getenv("part")
	if part == "part2" {
		part2(lines)
	} else {
		part1(lines)
	}
}
