package day24

import (
	"aoc-2023/aoc-lib"
	"os"
	"regexp"
)

func parseInput(f *os.File) []Hailstone {
	result := make([]Hailstone, 0)
	pattern := regexp.MustCompile(`(-?\d+), (-?\d+), (-?\d+) @ (-?\d+), (-?\d+), (-?\d+)`)
	for line := range aoc.LineReader(f) {
		match := pattern.FindStringSubmatch(line)
		result = append(result, Hailstone{
			position: FloatVec2{
				float64(aoc.MustParseInt(match[1])),
				float64(aoc.MustParseInt(match[2])),
			},
			velocity: FloatVec2{
				float64(aoc.MustParseInt(match[4])),
				float64(aoc.MustParseInt(match[5])),
			},
		})
	}
	return result
}
