package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

type Tree struct {
	height int32
	north  int32
	south  int32
	east   int32
	west   int32
}

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

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}

func max(a int32, b int32) int32 {
	if a >= b {
		return a
	}
	return b
}

func parseRows(lines []string, trees []Tree,
	columnCount int, fromRow int, toRow int) {

	defer waitGroup.Done()

	highest := '/'
	for row := fromRow; row < toRow; row++ {
		// Left to right
		for col := 0; col < columnCount; col++ {
			height := rune(lines[row][col])
			tree := &trees[(row*columnCount)+col]
			tree.height = height
			tree.west = highest
			highest = max(height, tree.west)
		}
		// right to left
		highest = '/'
		for col := columnCount - 1; col >= 0; col-- {
			height := rune(lines[row][col])
			tree := &trees[(row*columnCount)+col]
			tree.height = height
			tree.east = highest
			highest = max(height, tree.east)
		}
		highest = '/'
	}
}

func parseColumns(lines []string, trees []Tree,
	columnCount int, rowCount int, fromCol int, toCol int) {

	defer waitGroup.Done()

	highest := '/'
	for col := fromCol; col < toCol; col++ {
		// Top to bottom
		for row := 0; row < rowCount; row++ {
			height := rune(lines[row][col])
			// treeIndex := (row * columnCount) + col
			tree := &trees[(row*columnCount)+col]
			tree.height = height
			tree.north = highest
			highest = max(height, tree.north)
		}
		// Bottom to top
		highest = '/'
		for row := rowCount - 1; row >= 0; row-- {
			height := rune(lines[row][col])
			tree := &trees[(row*columnCount)+col]
			tree.height = height
			tree.south = highest
			highest = max(height, tree.south)
		}
		highest = '/'
	}
}

func countVisible(trees []Tree) int {
	visible := 0
	for _, tree := range trees {
		if tree.height > tree.north || tree.height > tree.south || tree.height > tree.east || tree.height > tree.west {
			visible++
		}
	}
	return visible
}

func part1(lines []string) {

	rowCount := len(lines)
	columnCount := len(lines[0])
	trees := make([]Tree, rowCount*columnCount)

	splitRow := rowCount / 2
	splitCol := columnCount / 2
	waitGroup.Add(4)
	go parseRows(lines, trees, columnCount, 0, splitRow)
	go parseRows(lines, trees, columnCount, splitRow, rowCount)
	go parseColumns(lines, trees, columnCount, rowCount, 0, splitCol)
	go parseColumns(lines, trees, columnCount, rowCount, splitCol, rowCount)

	// waitGroup.Add(2)
	//go parseRows(lines, trees, columnCount, 0, rowCount)
	//go parseColumns(lines, trees, columnCount, rowCount, 0, columnCount)

	waitGroup.Wait()

	visible := countVisible(trees)
	println(visible)

}

func part2(lines []string) {

	// for _, line := range lines {

	// }
	println("TODO")
}

func main() {
	// defer duration(track("main"))
	// lines := getInputAsArray("input-example.txt")
	lines := getInputAsArray("input.txt")

	part := os.Getenv("part")
	if part == "part2" {
		part2(lines)
	} else {
		part1(lines)
	}
}
