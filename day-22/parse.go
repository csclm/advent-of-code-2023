package day22

import (
	"aoc-2023/aoc-lib"
	"os"
	"regexp"
)

func parseInput(f *os.File) []Rect3 {
	result := make([]Rect3, 0)
	for line := range aoc.LineReader(f) {
		result = append(result, parseRect(line))
	}
	return result
}

func parseRect(line string) Rect3 {
	pattern := regexp.MustCompile(`(\d)+,(\d+),(\d+)~(\d+),(\d+),(\d+)`)
	match := pattern.FindStringSubmatch(line)
	startX := aoc.MustParseInt(match[1])
	startY := aoc.MustParseInt(match[2])
	startZ := aoc.MustParseInt(match[3])
	endX := aoc.MustParseInt(match[4]) + 1
	endY := aoc.MustParseInt(match[5]) + 1
	endZ := aoc.MustParseInt(match[6]) + 1
	return Rect3{
		x: aoc.Range{Min: startX, Max: endX},
		y: aoc.Range{Min: startY, Max: endY},
		z: aoc.Range{Min: startZ, Max: endZ},
	}
}
