package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/iochan"
)

func main() {
	f, _ := os.Open("./input.txt")
	numGrid := parseInput(f)
	graph := ClumsyCrucibleGraph{numGrid}

	// Way too slow, putting a pin in this one for now
	cost := graphSearch[ClumsyCrucibleNode](
		graph,
		ClumsyCrucibleNode{
			location:      Vec2{0, 0},
			lastDirection: Vec2{0, 0},
			momentum:      0,
		},
		func(ccn ClumsyCrucibleNode) bool {
			return ccn.location == Vec2{graph.Width() - 1, graph.Height() - 1}
		},
	)
	fmt.Printf("Minimum heat loss is %d\n", cost)
}

func parseInput(f *os.File) [][]int {
	result := make([][]int, 0)
	for line := range iochan.DelimReader(f, '\n') {
		numLine := make([]int, 0)
		for _, r := range strings.TrimSpace(line) {
			num, _ := strconv.ParseInt(string([]rune{r}), 10, 0)
			numLine = append(numLine, int(num))
		}
		result = append(result, numLine)
	}
	return result
}
