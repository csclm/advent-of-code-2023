package day3

import (
	"aoc-2023/aoc-lib"
	"os"
)

func readSchematicFromFile(f *os.File) Schematic {
	var contents [][]rune
	for schematicLine := range aoc.LineReader(f) {
		contents = append(contents, []rune(schematicLine))
	}
	return Schematic{
		contents: contents,
	}
}
