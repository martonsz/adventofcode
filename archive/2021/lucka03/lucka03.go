package lucka03

import (
	"adventofcode/aoc2021/lib"
	"fmt"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partA() {
	// lines := lib.GetInputAsArray("lucka03/input-example.txt")
	lines := lib.GetInputAsArray("lucka03/input.txt")

	lineCounter := 0
	var gammaSlice = make([]int, len(lines[0]))
	for _, line := range lines {
		lineCounter++
		//fmt.Printf("line: %v, chars: ", line)
		for pos, char := range line {
			bit, _ := strconv.Atoi(string(char))
			gammaSlice[pos] += bit
			//fmt.Printf("%v", bit)
		}
		//fmt.Printf(", %v\n", gammaSlice)
	}
	gamma := ""
	epsilon := ""
	for _, bit := range gammaSlice {

		if bit > (lineCounter / 2) {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	//fmt.Printf("gamma: %v, epsilon: %v\n", gamma, epsilon)
	//fmt.Printf("gamma: %v, epsilon: %v, multi: %v\n", gammaInt, epsilonInt, gammaInt*epsilonInt)
	fmt.Printf("partA: %v\n", gammaInt*epsilonInt)
}

func partB() {

}

func Solve() {
	fmt.Println("Lucka03")
	partA()
	partB()
}
