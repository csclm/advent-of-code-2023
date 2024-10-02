package day11

import (
	"aoc-2023/aoc-lib"
	"fmt"
	"os"
	"slices"

	"github.com/golang-collections/collections/set"
)

type Universe struct {
	width, height int
	galaxies      []aoc.Vec2
}

func Part1(f *os.File) {
	universe := parseGalaxies(f)
	universe.expand(2)
	fmt.Printf("Sum of path distances for 1x scale: %d\n", universe.sumOfPathDistancesBetweenGalaxies())
}

func Part2(f *os.File) {
	universe2 := parseGalaxies(f)
	universe2.expand(1000000)
	fmt.Printf("Sum of path distances for 1000000x scale: %d\n", universe2.sumOfPathDistancesBetweenGalaxies())
}

func (u *Universe) sumOfPathDistancesBetweenGalaxies() int {
	sumPathDistances := 0
	for i1 := 0; i1 < len(u.galaxies); i1++ {
		for i2 := 0; i2 < i1; i2++ {
			galaxy1 := u.galaxies[i1]
			galaxy2 := u.galaxies[i2]
			sumPathDistances += intAbs(galaxy1.X - galaxy2.X)
			sumPathDistances += intAbs(galaxy1.Y - galaxy2.Y)
		}
	}
	return sumPathDistances
}

func intAbs(in int) int {
	if in < 0 {
		return -in
	} else {
		return in
	}
}

func parseGalaxies(f *os.File) Universe {
	galaxies := make([]aoc.Vec2, 0)
	row := 0
	width := 0
	for line := range aoc.LineReader(f) {
		width = len(line)
		for col, r := range line {
			if r == '#' {
				galaxies = append(galaxies, aoc.NewVec2(row, col))
			}
		}
		row++
	}
	return Universe{width, row - 1, galaxies}
}

func (u *Universe) expand(scalingFactor int) {
	rowsWithGalaxies := set.New()
	colsWithGalaxies := set.New()
	for _, galaxy := range u.galaxies {
		rowsWithGalaxies.Insert(galaxy.Y)
		colsWithGalaxies.Insert(galaxy.X)
	}

	// Expand rows
	rowsToExpand := make([]int, 0)
	for i := 0; i < u.height; i++ {
		if !rowsWithGalaxies.Has(i) {
			rowsToExpand = append(rowsToExpand, i)
		}
	}
	slices.Sort(rowsToExpand)
	slices.Reverse(rowsToExpand)
	for _, rowToExpand := range rowsToExpand {
		for i, galaxy := range u.galaxies {
			if galaxy.Y > rowToExpand {
				u.galaxies[i].Y += scalingFactor - 1
			}
		}
	}

	// Expand columns
	colsToExpand := make([]int, 0)
	for i := 0; i < u.width; i++ {
		if !colsWithGalaxies.Has(i) {
			colsToExpand = append(colsToExpand, i)
		}
	}
	slices.Sort(colsToExpand)
	slices.Reverse(colsToExpand)
	for _, colToExpand := range colsToExpand {
		for i, galaxy := range u.galaxies {
			if galaxy.X > colToExpand {
				u.galaxies[i].X += scalingFactor - 1
			}
		}
	}
}
