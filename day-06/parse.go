package day6

import (
	"aoc-2023/aoc-lib"
	"os"
	"regexp"
	"strings"
)

func parseInput(f *os.File) []Race {
	lineReader := aoc.LineReader(f)
	timesLine := <-lineReader
	distancesLine := <-lineReader
	// Collapse spaces
	timesLine = regexp.MustCompile(` +`).ReplaceAllLiteralString(timesLine, " ")
	distancesLine = regexp.MustCompile(` +`).ReplaceAllLiteralString(distancesLine, " ")
	timesStrings := strings.Split(timesLine, " ")[1:]
	distancesStrings := strings.Split(distancesLine, " ")[1:]
	result := make([]Race, len(timesStrings))
	for i := 0; i < len(timesStrings); i++ {
		result = append(result, Race{
			time:     aoc.MustParseInt(timesStrings[i]),
			distance: aoc.MustParseInt(distancesStrings[i]),
		})
	}
	return result
}
