package lucka01

import (
	"bufio"
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

func partA() {
	f, err := os.Open("lucka01/input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	increase := 0
	previous := -1
	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		check(err)

		if previous != -1 && current > previous {
			increase++
		}
		previous = current
	}

	fmt.Printf("Part A: %v\n", increase)

	check(scanner.Err())
}

func partB() {

	//dat, _ := os.ReadFile("input-example.txt")
	dat, _ := os.ReadFile("lucka01/input.txt")

	increase := 0
	lines := strings.Split(string(dat), "\n")
	one, _ := strconv.Atoi(lines[0])
	two, _ := strconv.Atoi(lines[1])
	three, _ := strconv.Atoi(lines[2])
	prev := one + two + three

	for i := 3; i < len(lines); i++ {
		one, _ = strconv.Atoi(lines[i])
		two, _ = strconv.Atoi(lines[i-1])
		three, _ = strconv.Atoi(lines[i-2])
		current := one + two + three

		if current > prev {
			increase++
		}
		//fmt.Printf("%v  increase: %v\n", current, increase)
		prev = current
	}
	fmt.Printf("Part B: %v\n", increase)
}

func Solve() {
	fmt.Println("Lucka01")
	partA()
	partB()
}
