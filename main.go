package main

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	shades = []*color.Color{
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
	rowSize = len(shades)
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

func init() {
	for i := 0; i < rowSize; i++ {
		rootMap[i] = nil
	}
}

func main() {
	matrix := [][]int{}

	for i := 0; i < 100; i++ {
		matrix = append(matrix, createRow())
	}

	printMatrix(matrix)

}
