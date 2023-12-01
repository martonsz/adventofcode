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
func getInputAsArray(filePath string) []string {
	dat, err := os.ReadFile(filePath)
	check(err)
	return strings.Split(string(dat), "\n")
}

func part1(lines []string) {

	pairs := 0

	for _, line := range lines {

		pairSplit := strings.Split(line, ",")

		elf1 := strings.Split(pairSplit[0], "-")
		elf2 := strings.Split(pairSplit[1], "-")

		elf1Start, _ := strconv.Atoi(elf1[0])
		elf1End, _ := strconv.Atoi(string(elf1[1]))

		elf2Start, _ := strconv.Atoi(elf2[0])
		elf2End, _ := strconv.Atoi(string(elf2[1]))

		if (elf1Start >= elf2Start && elf1End <= elf2End) || (elf2Start >= elf1Start && elf2End <= elf1End) {
			pairs++
		}

	}
	fmt.Println(pairs)
}

func part2(lines []string) {
	pairs := 0

	for _, line := range lines {

		pairSplit := strings.Split(line, ",")

		elf1 := strings.Split(pairSplit[0], "-")
		elf2 := strings.Split(pairSplit[1], "-")

		elf1Start, _ := strconv.Atoi(elf1[0])
		elf1End, _ := strconv.Atoi(string(elf1[1]))

		elf2Start, _ := strconv.Atoi(elf2[0])
		elf2End, _ := strconv.Atoi(string(elf2[1]))

		if elf2Start <= elf1End && elf2End >= elf1Start {
			pairs++
		}

	}
	fmt.Println(pairs)
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
