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

func printCWD() {
	os.Chdir("/Users")
	newDir, _ := os.Getwd()
	fmt.Printf("Current Working Direcotry: %s\n", newDir)
}

func part1(lines []string) {

	highestCaloriesCount := 0
	caloriesCount := 0
	for _, line := range lines {

		calories, err := strconv.Atoi(line)

		if err == nil {
			caloriesCount += calories
		} else {
			if caloriesCount > highestCaloriesCount {
				highestCaloriesCount = caloriesCount
			}
			caloriesCount = 0
		}
	}
	fmt.Println(highestCaloriesCount)
}

func part2(lines []string) {

	highestCaloriesCount1 := 0
	highestCaloriesCount2 := 0
	highestCaloriesCount3 := 0
	caloriesCount := 0
	for _, line := range lines {

		calories, err := strconv.Atoi(line)

		if err == nil {
			caloriesCount += calories
		} else {
			if caloriesCount > highestCaloriesCount1 {
				highestCaloriesCount3 = highestCaloriesCount2
				highestCaloriesCount2 = highestCaloriesCount1
				highestCaloriesCount1 = caloriesCount
			} else if caloriesCount > highestCaloriesCount2 {
				highestCaloriesCount3 = highestCaloriesCount2
				highestCaloriesCount2 = caloriesCount
			} else if caloriesCount > highestCaloriesCount3 {
				highestCaloriesCount3 = caloriesCount
			}
			caloriesCount = 0
		}
	}
	if caloriesCount > 0 {
		if caloriesCount > highestCaloriesCount1 {
			highestCaloriesCount3 = highestCaloriesCount2
			highestCaloriesCount2 = highestCaloriesCount1
			highestCaloriesCount1 = caloriesCount
		} else if caloriesCount > highestCaloriesCount2 {
			highestCaloriesCount3 = highestCaloriesCount2
			highestCaloriesCount2 = caloriesCount
		} else if caloriesCount > highestCaloriesCount3 {
			highestCaloriesCount3 = caloriesCount
		}
		caloriesCount = 0
	}
	fmt.Println(highestCaloriesCount1 + highestCaloriesCount2 + highestCaloriesCount3)
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
