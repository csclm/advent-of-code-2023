package day9

import (
	"aoc-2023/aoc-lib"
	"os"
)

func parseInputFile(f *os.File) [][]int {
	sequences := make([][]int, 0)
	for line := range aoc.LineReader(f) {
		nums := aoc.MustParseListOfNums(line, " ")
		sequences = append(sequences, nums)
	}
	return sequences
}
