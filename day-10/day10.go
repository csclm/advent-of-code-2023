package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./input.txt")
	grid := parseInput(f)

	part1Answer := CalculateFurthestStep(grid)
	fmt.Printf("Number of steps for furthest distance: %d\n", part1Answer)

	enclosedArea := CalculateEnclosedArea(&grid)
	fmt.Printf("Enclosed area: %d\n", enclosedArea)
}
