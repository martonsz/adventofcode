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

func part1(lines []string) {

	for _, line := range lines {

	}
	fmt.Println("TODO")
}

func part2(lines []string) {

	for _, line := range lines {

	}
	println("TODO")
}

var prioMap = make(map[int32]int)

func main() {

	lines := getInputAsArray("input.txt")

	part := os.Getenv("part")
	if part == "part2" {
		part2(lines)
	} else {
		part1(lines)
	}
}
