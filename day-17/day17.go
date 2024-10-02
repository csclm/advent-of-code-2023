package day17

import (
	"aoc-2023/aoc-lib"
	"fmt"
	"os"
)

func Part1(f *os.File) {
	numGrid := parseInput(f)
	graphPart1 := ClumsyCrucibleGraph{contents: numGrid, maxStraightLine: 3, minStraightLine: 0}
	costPart1 := graphSearch[ClumsyCrucibleNode](
		graphPart1,
		CrucibleStartingNode(),
		IsFinishingNode(graphPart1),
	)
	fmt.Printf("Part 1 minimum heat loss is %d\n", costPart1)
}

func Part2(f *os.File) {
	numGrid := parseInput(f)
	graphPart2 := ClumsyCrucibleGraph{contents: numGrid, maxStraightLine: 10, minStraightLine: 4}
	costPart2 := graphSearch[ClumsyCrucibleNode](
		graphPart2,
		CrucibleStartingNode(),
		IsFinishingNode(graphPart2),
	)
	fmt.Printf("Part 2 minimum heat loss is %d\n", costPart2)
}

func CrucibleStartingNode() ClumsyCrucibleNode {
	return ClumsyCrucibleNode{
		location:      aoc.NewVec2(0, 0),
		lastDirection: aoc.NewVec2(0, 0),
		momentum:      0,
	}
}
func IsFinishingNode(graph ClumsyCrucibleGraph) func(ClumsyCrucibleNode) bool {
	return func(node ClumsyCrucibleNode) bool {
		return node.location == aoc.NewVec2(graph.Width()-1, graph.Height()-1)
	}
}

func parseInput(f *os.File) [][]int {
	result := make([][]int, 0)
	for line := range aoc.LineReader(f) {
		numLine := make([]int, 0)
		for _, r := range line {
			numLine = append(numLine, aoc.MustParseDigit(r))
		}
		result = append(result, numLine)
	}
	return result
}
