package main

import (
	"container/list"
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

func getStacks(lines []string) []list.List {
	var stackCount int
	for _, line := range lines {
		if !strings.Contains(line, "[") {
			line = strings.Trim(line, " ")
			stackCount, _ = strconv.Atoi(string(line[len(line)-1]))
			break
		}
	}
	stacks := make([]list.List, stackCount)
	return stacks
}

func solve(lines []string, part int) {

	stacks := getStacks(lines)

	for _, line := range lines {

		if strings.Contains(line, "[") {
			stackIndex := -1
			for i := 1; i < len(line); i += 4 {
				container := line[i]
				stackIndex++
				if container == ' ' {
					continue
				}
				stacks[stackIndex].PushBack(string(container))
			}
		} else if strings.HasPrefix(line, "move") {
			splitLine := strings.Split(line, " ")
			count, _ := strconv.Atoi(splitLine[1])
			from, _ := strconv.Atoi(splitLine[3])
			to, _ := strconv.Atoi(splitLine[5])

			pickStack := list.New()
			for i := 0; i < count; i++ {
				containerEl := stacks[from-1].Front()
				stacks[from-1].Remove(containerEl)
				if part == 1 {
					stacks[to-1].PushFront(containerEl.Value)
				} else {
					pickStack.PushBack(containerEl.Value)
				}
			}
			if part == 2 {
				stacks[to-1].PushFrontList(pickStack)
			}
		}
	}

	for _, l := range stacks {
		fmt.Printf("%s", l.Front().Value)
	}
}

func main() {

	lines := getInputAsArray("input.txt")

	part := os.Getenv("part")
	if part == "part2" {
		solve(lines, 2)
	} else {
		solve(lines, 1)
	}
}
