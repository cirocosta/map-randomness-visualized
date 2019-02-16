package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"

	"github.com/fatih/color"
)

var (
	shouldShuffle = flag.Bool("shuffle", false, "whether we should manually shuffle or not")
	shades        = []*color.Color{
		color.New(color.FgBlack),
		color.New(color.FgRed),
		color.New(color.FgGreen),
		color.New(color.FgYellow),
		color.New(color.FgBlue),
		color.New(color.FgMagenta),
		color.New(color.FgCyan),
		color.New(color.FgWhite),
		color.New(color.FgHiBlack),
		color.New(color.FgHiRed),
		color.New(color.FgHiGreen),
		color.New(color.FgHiYellow),
		color.New(color.FgHiBlue),
		color.New(color.FgHiMagenta),
		color.New(color.FgHiCyan),
		color.New(color.FgHiWhite),
	}
	rowSize = 2
	rootMap = map[int]interface{}{}
)

func printMatrix(matrix [][]int) {
	for _, i := range matrix {
		for _, j := range i {
			shades[j].Print("█")
		}
		fmt.Print("\n")
	}
}

func createRow() []int {
	var (
		actualMap = map[int]interface{}{}
		row       = make([]int, rowSize)
	)

	for k, v := range rootMap {
		actualMap[k] = v
	}

	idx := 0
	for k := range actualMap {
		row[idx] = k
		idx++
	}

	return row
}

func shuffle(v sort.Interface) {
	for i := v.Len() - 1; i > 0; i-- {
		v.Swap(i, rand.Intn(i+1))
	}
}

func createShuffledRow() []int {
	baseRow := []int{}
	for k := range rootMap {
		baseRow = append(baseRow, k)
	}

	shuffled := make([]int, len(baseRow))
	copy(shuffled, baseRow)
	shuffle(sort.IntSlice(shuffled))

	return shuffled
}

func init() {
	for i := 0; i < rowSize; i++ {
		rootMap[i] = nil
	}
}

var rowCreator func() []int

func main() {
	flag.Parse()

	if *shouldShuffle {
		rowCreator = createShuffledRow
	} else {
		rowCreator = createRow
	}

	matrix := [][]int{}

	for i := 0; i < 100; i++ {
		matrix = append(matrix, rowCreator())
	}

	printMatrix(matrix)

}
