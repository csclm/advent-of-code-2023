package main

import (
	"os"
	"strings"

	"github.com/mitchellh/iochan"
)

func readSchematicFromFile(f *os.File) Schematic {
	var contents [][]rune
	for line := range iochan.DelimReader(f, '\n') {
		schematicLine := strings.TrimSpace(line)
		contents = append(contents, []rune(schematicLine))
	}
	return Schematic{
		contents: contents,
	}
}
