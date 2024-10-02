package day13

import (
	"aoc-2023/aoc-lib"
	"os"
)

func parseInput(f *os.File) []Grid {
	grids := make([]Grid, 0)
	currentGridContents := make([][]rune, 0)
	for line := range aoc.LineReader(f) {
		if len(line) == 0 {
			grids = append(grids, Grid{currentGridContents, false})
			currentGridContents = make([][]rune, 0)
		} else {
			currentGridContents = append(currentGridContents, []rune(line))
		}
	}
	grids = append(grids, Grid{currentGridContents, false})
	return grids
}
