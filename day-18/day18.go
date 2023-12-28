package day18

import (
	"aoc-2023/aoc-lib"
	"fmt"
	"os"
)

func Part1(f *os.File) {
	instructions := parseInput(f)
	holes := digHoles(instructions)
	ground := makeGridFromHoles(holes)
	digInteriorHoles(ground)

	totalDug := 0
	width := len(ground[0])
	height := len(ground)
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if ground[row][col].dug {
				totalDug++
			}
		}
	}

	fmt.Printf("Part 1 Total Volume: %d\n", totalDug)
}

func Part2(f *os.File) {
	instructions := parseInputWithHexInstructions(f)
	vertices := verticesFromDigInstructions(instructions)
	area := aoc.IntAbs(shoelace(vertices))
	runTestCases()
	fmt.Printf("Part 2 Total Volume: %d\n", area)
}
