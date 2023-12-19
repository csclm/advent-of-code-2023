package main

import (
	"fmt"
	"os"
)

type Grid struct {
	contents  [][]rune
	transpose bool
}

func (g Grid) RuneAt(row int, col int) rune {
	if g.transpose {
		return g.contents[col][row]
	} else {
		return g.contents[row][col]
	}
}

func (g Grid) Transposed() Grid {
	return Grid{
		g.contents,
		(!g.transpose),
	}
}

func (g Grid) Width() int {
	if g.transpose {
		return len(g.contents)
	} else {
		return len(g.contents[0])
	}
}

func (g Grid) Height() int {
	if g.transpose {
		return len(g.contents[0])
	} else {
		return len(g.contents)
	}
}

func (g Grid) Print() {
	for _, row := range g.contents {
		fmt.Println(string(row))
	}
}

func main() {
	f, _ := os.Open("./input.txt")
	grids := parseInput(f)

	summary := 0
	for _, grid := range grids {
		vertical := findReflectionColumns(grid)
		for _, line := range vertical {
			summary += (line + 1)
		}
		horizontal := findReflectionColumns(grid.Transposed())
		for _, line := range horizontal {
			summary += (line + 1) * 100
		}
	}

	summarySmudged := 0
	for _, grid := range grids {
		vertical := findReflectionColumnsWithSmudge(grid)
		for _, line := range vertical {
			summarySmudged += (line + 1)
		}
		horizontal := findReflectionColumnsWithSmudge(grid.Transposed())
		for _, line := range horizontal {
			summarySmudged += (line + 1) * 100
		}
	}

	fmt.Printf("Summary: %d\n", summary)
	fmt.Printf("Summary with smudges: %d\n", summarySmudged)
}
