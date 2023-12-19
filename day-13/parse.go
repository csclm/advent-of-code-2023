package main

import (
	"os"
	"strings"

	"github.com/mitchellh/iochan"
)

func parseInput(f *os.File) []Grid {
	grids := make([]Grid, 0)
	currentGridContents := make([][]rune, 0)
	for line := range iochan.DelimReader(f, '\n') {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) == 0 {
			grids = append(grids, Grid{currentGridContents, false})
			currentGridContents = make([][]rune, 0)
		} else {
			currentGridContents = append(currentGridContents, []rune(trimmedLine))
		}
	}
	grids = append(grids, Grid{currentGridContents, false})
	return grids
}
