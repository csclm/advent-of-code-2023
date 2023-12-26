package day10

import (
	"fmt"
	"os"
)

func Part1(f *os.File) {
	grid := parseInput(f)
	part1Answer := CalculateFurthestStep(grid)
	fmt.Printf("Number of steps for furthest distance: %d\n", part1Answer)
}

func Part2(f *os.File) {
	grid := parseInput(f)
	enclosedArea := CalculateEnclosedArea(&grid)
	fmt.Printf("Enclosed area: %d\n", enclosedArea)
}
