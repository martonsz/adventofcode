package main

import (
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

func part1(lines []string) {

	line := lines[0]
	for i := 0; i < len(line); i++ {
		set := make(map[byte]bool)
		if i >= 3 {
			set[line[i]] = true
			set[line[i-1]] = true
			set[line[i-2]] = true
			set[line[i-3]] = true
			if len(set) == 4 {
				println(i + 1)
				break
			}
		}
	}
}

func part2(lines []string) {
	line := lines[0]
	for i := 0; i < len(line); i++ {
		set := make(map[byte]bool)
		if i >= 13 {
			for z := 0; z < 14; z++ {
				set[line[i-z]] = true
			}
			if len(set) == 14 {
				println(i + 1)
				break
			}
		}
	}
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
