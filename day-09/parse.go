package day9

import (
	"aoc-2023/aoc-lib"
	"os"
	"strconv"
	"strings"
)

func parseInputFile(f *os.File) [][]int {
	sequences := make([][]int, 0)
	for line := range aoc.LineReader(f) {
		numStrings := strings.Split(line, " ")
		nums := make([]int, 0)
		for _, numStr := range numStrings {
			num, _ := strconv.ParseInt(numStr, 10, 0)
			nums = append(nums, int(num))
		}
		sequences = append(sequences, nums)
	}
	return sequences
}
