package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/golang-collections/collections/set"
	"github.com/mitchellh/iochan"
)

type Vec2 struct {
	x, y int
}

type Universe struct {
	width, height int
	galaxies      []Vec2
}

func main() {
	f, _ := os.Open("./input.txt")
	universe := parseGalaxies(f)
	universe.expand(2)
	fmt.Printf("Sum of path distances for 1x scale: %d\n", universe.sumOfPathDistancesBetweenGalaxies())

	f.Seek(0, 0)
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
			sumPathDistances += intAbs(galaxy1.x - galaxy2.x)
			sumPathDistances += intAbs(galaxy1.y - galaxy2.y)
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
	galaxies := make([]Vec2, 0)
	row := 0
	width := 0
	for line := range iochan.DelimReader(f, '\n') {
		width = len(strings.TrimSpace(line))
		for col, r := range line {
			if r == '#' {
				galaxies = append(galaxies, Vec2{row, col})
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
		rowsWithGalaxies.Insert(galaxy.y)
		colsWithGalaxies.Insert(galaxy.x)
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
			if galaxy.y > rowToExpand {
				u.galaxies[i].y += scalingFactor - 1
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
			if galaxy.x > colToExpand {
				u.galaxies[i].x += scalingFactor - 1
			}
		}
	}
}
