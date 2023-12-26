package day14

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/iochan"
)

func Part1(f *os.File) {
	grid := parseInput(f)
	slideStonesNorth(grid)
	totalLoad := 0
	for ri := 0; ri < grid.Height(); ri++ {
		for ci := 0; ci < grid.Width(); ci++ {
			if grid.RuneAt(ri, ci) == RoundStone {
				totalLoad += grid.Height() - ri
			}
		}
	}
	fmt.Printf("Total load: %d\n", totalLoad)
}

const RoundStone = 'O'
const SquareStone = '#'
const EmptySpace = '.'

func parseInput(f *os.File) Grid {
	result := make([][]rune, 0)
	for line := range iochan.DelimReader(f, '\n') {
		result = append(result, []rune(strings.TrimSpace(line)))
	}
	return Grid{result, North}
}
