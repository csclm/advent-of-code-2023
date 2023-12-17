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
	universe.expand(1)
	fmt.Printf("Sum of path distances for 1x scale: %d\n", universe.sumOfPathDistancesBetweenGalaxies())

	f.Seek(0, 0)
	universe2 := parseGalaxies(f)
	universe2.expand(1000000)
	fmt.Printf("Sum of path distances for 1000000x scale: %d\n", universe2.sumOfPathDistancesBetweenGalaxies())
}

func (u *Universe) sumOfPathDistancesBetweenGalaxies() int {
	sumPathDistances := 0
	for pair := range unorderedPairs(len(u.galaxies)) {
		galaxy1 := u.galaxies[pair.x]
		galaxy2 := u.galaxies[pair.y]
		sumPathDistances += intAbs(galaxy1.x - galaxy2.x)
		sumPathDistances += intAbs(galaxy1.y - galaxy2.y)
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

func unorderedPairs(span int) chan Vec2 {
	result := make(chan Vec2)
	go func() {
		for i1 := 0; i1 < span; i1++ {
			for i2 := 0; i2 < i1; i2++ {
				result <- Vec2{i1, i2}
			}
		}
		close(result)
	}()
	return result
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
	rowsToExpand := u.rowsWithoutGalaxies()
	slices.Sort(rowsToExpand)
	slices.Reverse(rowsToExpand)
	for _, rowToExpand := range rowsToExpand {
		for i, galaxy := range u.galaxies {
			if galaxy.y > rowToExpand {
				u.galaxies[i].y += scalingFactor
			}
		}
	}
	colsToExpand := u.colsWithoutGalaxies()
	slices.Sort(colsToExpand)
	slices.Reverse(colsToExpand)
	for _, colToExpand := range colsToExpand {
		for i, galaxy := range u.galaxies {
			if galaxy.x > colToExpand {
				u.galaxies[i].x += scalingFactor
			}
		}
	}
}

func (u *Universe) rowsWithoutGalaxies() []int {
	allRows := set.New()
	for row := 0; row < u.height; row++ {
		allRows.Insert(row)
	}
	rowsWithGalaxies := set.New()
	for _, galaxy := range u.galaxies {
		rowsWithGalaxies.Insert(galaxy.y)
	}
	rowsWithoutGalaxiesSet := allRows.Difference(rowsWithGalaxies)
	result := make([]int, 0)
	rowsWithoutGalaxiesSet.Do(func(row interface{}) {
		result = append(result, row.(int))
	})
	return result
}

func (u *Universe) colsWithoutGalaxies() []int {
	allCols := set.New()
	for col := 0; col < u.width; col++ {
		allCols.Insert(col)
	}
	colsWithGalaxies := set.New()
	for _, galaxy := range u.galaxies {
		colsWithGalaxies.Insert(galaxy.x)
	}
	colsWithoutGalaxiesSet := allCols.Difference(colsWithGalaxies)
	result := make([]int, 0)
	colsWithoutGalaxiesSet.Do(func(col interface{}) {
		result = append(result, col.(int))
	})
	return result
}
