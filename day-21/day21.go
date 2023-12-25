package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-collections/collections/set"
	"github.com/mitchellh/iochan"
)

func main() {
	f, _ := os.Open("./input.txt")
	grid := parseInput(f)

	currentLocations := set.New()
	start := findOccurrences(grid, 'S')[0]
	currentLocations.Insert(start)

	replaceOccurrences(grid, 'S', '.')

	for i := 0; i < 64; i++ {
		newLocations := set.New()
		currentLocations.Do(func(i interface{}) {
			location := i.(Vec2)
			for _, dir := range cardinalDirections() {
				newLocation := location.Plus(dir)
				if grid[newLocation.y][newLocation.x] == '.' {
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

func findOccurrences(grid [][]rune, runeToFind rune) []Vec2 {
	result := make([]Vec2, 0)
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == runeToFind {
				result = append(result, Vec2{col, row})
			}
		}
	}
	return result
}

func parseInput(f *os.File) [][]rune {
	result := make([][]rune, 0)
	for line := range iochan.DelimReader(f, '\n') {
		result = append(result, []rune(strings.TrimSpace(line)))
	}
	return result
}
