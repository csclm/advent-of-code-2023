package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./input.txt")
	network, directions := parseInput(f)

	part1TraversalSteps := numberOfTraversalStepsPart1(directions, network)
	fmt.Printf("Number of steps from AAA to ZZZ: %d\n", part1TraversalSteps)

	part2TraversalSteps := numberOfTraversalStepsPart2(directions, network)
	fmt.Printf("Number of steps from **A to **Z: %d\n", part2TraversalSteps)
}
