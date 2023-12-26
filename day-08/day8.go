package day8

import (
	"fmt"
	"os"
)

func Part1(f *os.File) {
	network, directions := parseInput(f)

	part1TraversalSteps := numberOfTraversalStepsPart1(directions, network)
	fmt.Printf("Number of steps from AAA to ZZZ: %d\n", part1TraversalSteps)
}

func Part2(f *os.File) {
	network, directions := parseInput(f)
	part2TraversalSteps := numberOfTraversalStepsPart2(directions, network)
	fmt.Printf("Number of steps from **A to **Z: %d\n", part2TraversalSteps)
}
