package main

import (
	"fmt"
	"os"
	"sort"
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

func parse(trees []Tree, lines []string, rowCount int, columnCount int) {

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

func countScenicScore(trees []Tree, scenicScores []int64, rowCount int, columnCount int, row int, col int) {
	// defer waitGroup.Done()

	index := (row * columnCount) + col
	baseHeight := trees[index].height
	var scenicScore int64 = 1

	// <--
	viewLength := 0
	for x := col - 1; x >= 0; x-- {
		height := trees[(row*columnCount)+x].height
		viewLength++
		if baseHeight <= height {
			break
		}
	}
	scenicScore *= int64(viewLength)

	// -->
	viewLength = 0
	for x := col + 1; x < columnCount; x++ {
		height := trees[(row*columnCount)+x].height
		viewLength++
		if baseHeight <= height {
			break
		}
	}
	scenicScore *= int64(viewLength)

	// up
	viewLength = 0
	for y := row - 1; y >= 0; y-- {
		height := trees[(y*columnCount)+col].height
		viewLength++
		if baseHeight <= height {
			break
		}
	}
	scenicScore *= int64(viewLength)

	// down
	viewLength = 0
	for y := row + 1; y < rowCount; y++ {
		height := trees[(y*columnCount)+col].height
		viewLength++
		if baseHeight <= height {
			break
		}
	}
	scenicScore *= int64(viewLength)
	scenicScores[index] = scenicScore
}

func getBestScenicScore(trees []Tree, rowCount int, columnCount int) int64 {

	// waitGroup.Add(rowCount * columnCount)
	scenicScores := make([]int64, rowCount*columnCount)
	for row := 0; row < rowCount; row++ {
		for col := 0; col < columnCount; col++ {
			countScenicScore(trees, scenicScores, rowCount, columnCount, row, col)
		}
	}
	// waitGroup.Wait()

	sort.Slice(scenicScores, func(i, j int) bool { return scenicScores[i] > scenicScores[j] })
	return scenicScores[0]
}

func main() {
	// defer duration(track("main"))

	// lines := getInputAsArray("input-example.txt")
	lines := getInputAsArray("input.txt")

	rowCount := len(lines)
	columnCount := len(lines[0])
	trees := make([]Tree, rowCount*columnCount)
	parse(trees, lines, rowCount, columnCount)

	part := os.Getenv("part")
	if part == "part2" {
		bestScenicScore := getBestScenicScore(trees, rowCount, columnCount)
		println(bestScenicScore)
	} else {
		visible := countVisible(trees)
		println(visible)
	}
}
