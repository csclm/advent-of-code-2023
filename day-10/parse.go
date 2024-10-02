package day10

import (
	"aoc-2023/aoc-lib"
	"os"
)

func parseInput(f *os.File) PipeGrid {
	result := make([][]PipeRune, 0)
	for line := range aoc.LineReader(f) {
		result = append(result, []PipeRune(line))
	}
	return PipeGrid(result)
}
