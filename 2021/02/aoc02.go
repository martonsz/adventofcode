package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var regexpSpace = regexp.MustCompile("\\s")

type action struct {
	horizontal int
	depth      int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLine(line string) action {
	split := regexpSpace.Split(line, -1)
	a := action{}
	switch split[0] {
	case "forward":
		a.horizontal, _ = strconv.Atoi(split[1])
		a.depth = 0
	case "up":
		a.horizontal = 0
		a.depth, _ = strconv.Atoi(split[1])
		a.depth *= -1
	case "down":
		a.horizontal = 0
		a.depth, _ = strconv.Atoi(split[1])
	}

	return a
}

func partA() {
	//f, _ := os.Open("02/input-example.txt")
	f, _ := os.Open("02/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	horizontal := 0
	depth := 0
	for scanner.Scan() {
		a := parseLine(scanner.Text())

		horizontal += a.horizontal
		depth += a.depth
		fmt.Printf("%v %v %v, %v\n", a, horizontal, depth, horizontal*depth)
	}

	check(scanner.Err())
}

func partB() {

}

func main() {
	partA()
	partB()
}
