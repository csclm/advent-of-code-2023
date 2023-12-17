package main

import (
	"os"
	"strings"

	"github.com/mitchellh/iochan"
)

func parseInput(f *os.File) PipeGrid {
	result := make([][]PipeRune, 0)
	for line := range iochan.DelimReader(f, '\n') {
		result = append(result, []PipeRune(strings.TrimSpace(line)))
	}
	return PipeGrid(result)
}
