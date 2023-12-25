package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./input.txt")
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

	// f.Seek(0, 0)
	// part2Instructions := parseInputWithHexInstructions(f)
	/*
		Part 2 idea:
		parse each set of 2 consecutive moves as a rectangle
		for any pair of intersecting rectangles, re-partition them as a set of non-intersecting rectangles
		.. that might be O(n^2) or O(n^3) so maybe a quadtree would speed it up?
	*/
}
