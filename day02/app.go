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

	moveMap := make(map[byte]int)
	moveMap['A'] = 0
	moveMap['B'] = 1
	moveMap['C'] = 2
	moveMap['X'] = 0
	moveMap['Y'] = 1
	moveMap['Z'] = 2

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
	moveMap := make(map[string]int)
	moveMap["A X"] = 3
	moveMap["A Y"] = 4
	moveMap["A Z"] = 8
	moveMap["B X"] = 1
	moveMap["B Y"] = 5
	moveMap["B Z"] = 9
	moveMap["C X"] = 2
	moveMap["C Y"] = 6
	moveMap["C Z"] = 7

	totalScore := 0
	for _, line := range lines {
		totalScore += moveMap[line]
	}
	fmt.Println(totalScore)
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
