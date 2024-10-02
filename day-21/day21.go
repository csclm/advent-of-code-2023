package day21

import (
	"aoc-2023/aoc-lib"
	"fmt"
	"os"

	"github.com/golang-collections/collections/set"
)

func Part1(f *os.File) {
	grid := parseInput(f)

	currentLocations := set.New()
	start := findOccurrences(grid, 'S')[0]
	currentLocations.Insert(start)

	replaceOccurrences(grid, 'S', '.')

	for i := 0; i < 64; i++ {
		newLocations := set.New()
		currentLocations.Do(func(i interface{}) {
			location := i.(aoc.Vec2)
			for _, dir := range aoc.CardinalDirections() {
				newLocation := location.Plus(dir)
				if grid[newLocation.Y][newLocation.X] == '.' {
					newLocations.Insert(newLocation)
				}
			}
		})
		currentLocations = newLocations
	}

	fmt.Printf("Garden plots reachable in 64 steps: %d\n", currentLocations.Len())
}

func replaceOccurrences(grid [][]rune, runeToFind rune, runeToReplaceWith rune) {
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == runeToFind {
				grid[row][col] = runeToReplaceWith
			}
		}
	}
}

func findOccurrences(grid [][]rune, runeToFind rune) []aoc.Vec2 {
	result := make([]aoc.Vec2, 0)
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == runeToFind {
				result = append(result, aoc.NewVec2(col, row))
			}
		}
	}
	return result
}

func parseInput(f *os.File) [][]rune {
	result := make([][]rune, 0)
	for line := range aoc.LineReader(f) {
		result = append(result, []rune(line))
	}
	return result
}
