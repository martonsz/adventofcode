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
func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func part2(lines []string) {
	sum := 0

	for i := 0; i < len(lines); i += 3 {
		elf1Items := lines[i]
		elf2Items := lines[i+1]
		elf3Items := lines[i+2]
		for item, prio := range prioMap {
			if strings.ContainsRune(elf1Items, item) &&
				strings.ContainsRune(elf2Items, item) &&
				strings.ContainsRune(elf3Items, item) {
				sum += prio
			}
		}
	}
	println(sum)
}

var prioMap = make(map[int32]int)

func main() {

	aIntValue := int('a') - 1
	for rune := 'a'; rune <= 'z'; rune++ {
		prioMap[rune] = int(rune) - aIntValue
	}
	AIntValue := int('A') - prioMap['z'] - 1
	for rune := 'A'; rune <= 'Z'; rune++ {
		prioMap[rune] = int(rune) - AIntValue
	}

	lines := getInputAsArray("input.txt")

	part := os.Getenv("part")
	if part == "part2" {
		part2(lines)
	} else {
		part1(lines)
	}
}
