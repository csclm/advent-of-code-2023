package day9

import (
	"aoc-2023/aoc-lib"
	"os"
	"strings"
)

func parseInputFile(f *os.File) [][]int {
	sequences := make([][]int, 0)
	for line := range aoc.LineReader(f) {
		numStrings := strings.Split(line, " ")
		nums := make([]int, 0)
		for _, numStr := range numStrings {
			nums = append(nums, aoc.MustParseInt(numStr))
		}
		sequences = append(sequences, nums)
	}
	return sequences
}
