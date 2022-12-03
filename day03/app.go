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

	prioMap := make(map[int32]int)

	aIntValue := int('a') - 1
	for rune := 'a'; rune <= 'z'; rune++ {
		prioMap[rune] = int(rune) - aIntValue
	}
	AIntValue := int('A') - prioMap['z'] - 1
	for rune := 'A'; rune <= 'Z'; rune++ {
		prioMap[rune] = int(rune) - AIntValue
	}

	sum := 0
	for _, line := range lines {
		length := len(line)
		compartmentA := line[:length/2]
		compartmentB := line[length/2:]
		for _, item := range compartmentA {
			if strings.ContainsRune(compartmentB, item) {
				sum += prioMap[item]
				break
			}
		}
	}
	fmt.Println(sum)
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
