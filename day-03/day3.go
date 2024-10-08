package day3

import (
	"aoc-2023/aoc-lib"
	"fmt"
	"os"
	"regexp"
)

type Schematic struct {
	contents [][]rune
}

type PartNumberLocation struct {
	row    int
	column int
	length int
}

// rune, is-in-bounds
func (schematic *Schematic) RuneAt(row int, column int) (rune, bool) {
	if column >= 0 && column < len(schematic.contents[0]) && row >= 0 && row < len(schematic.contents) {
		return schematic.contents[row][column], true
	} else {
		return '\x00', false
	}
}

func locateNumbers(schematic *Schematic) []PartNumberLocation {
	pattern := regexp.MustCompile("\\d+")
	result := make([]PartNumberLocation, 0)
	for rowNum, row := range schematic.contents {
		for _, match := range pattern.FindAllStringIndex(string(row), -1) {
			result = append(result, PartNumberLocation{
				row:    rowNum,
				column: match[0],
				length: match[1] - match[0],
			})
		}
	}
	return result
}

func (schematic *Schematic) isNumberAPart(location PartNumberLocation) bool {
	for i := -1; i <= location.length; i++ {
		up, upInBounds := schematic.RuneAt(location.row-1, location.column+i)
		if upInBounds && isSchematicSymbol(up) {
			return true
		}
		middle, middleInBounds := schematic.RuneAt(location.row, location.column+i)
		if middleInBounds && isSchematicSymbol(middle) {
			return true
		}
		down, downInBounds := schematic.RuneAt(location.row+1, location.column+i)
		if downInBounds && isSchematicSymbol(down) {
			return true
		}
	}
	return false
}

func getGearingRatio(schematic *Schematic, partNumberMap *[][]int, row int, col int) int {
	centerRune, centerInBounds := schematic.RuneAt(row, col)
	if !centerInBounds || centerRune != '*' {
		return 0
	}
	topParts, topRatio := partsInTriple(partNumberMap, row-1, col)
	middleParts, middleRatio := partsInTriple(partNumberMap, row, col)
	bottomParts, bottomRatio := partsInTriple(partNumberMap, row+1, col)
	if topParts+middleParts+bottomParts != 2 {
		return 0
	}
	return topRatio * middleRatio * bottomRatio
}

// number of parts, ratio of the part numbers
func partsInTriple(partNumberMap *[][]int, row int, centerCol int) (int, int) {
	centerPartNumber := (*partNumberMap)[row][centerCol]
	if centerPartNumber != 0 {
		// Doesn't matter about left and right - they'd be part of the same part
		return 1, centerPartNumber
	}
	parts := 0
	ratio := 1
	leftPartNumber := (*partNumberMap)[row][centerCol-1]
	rightPartNumber := (*partNumberMap)[row][centerCol+1]
	if leftPartNumber != 0 {
		parts++
		ratio *= leftPartNumber
	}
	if rightPartNumber != 0 {
		parts++
		ratio *= rightPartNumber
	}
	return parts, ratio
}

func generatePartNumberMap(schematic *Schematic, locations *[]PartNumberLocation) [][]int {
	result := make([][]int, len(schematic.contents))
	for i := range schematic.contents {
		result[i] = make([]int, len(schematic.contents[i]))
	}

	for _, location := range *locations {
		partNumber := schematic.getNumberAtLocation(location)
		for i := location.column; i < location.column+location.length; i++ {
			result[location.row][i] = partNumber
		}
	}

	return result
}

func (schematic *Schematic) getNumberAtLocation(location PartNumberLocation) int {
	numAsString := string(schematic.contents[location.row][location.column : location.column+location.length])
	return aoc.MustParseInt(numAsString)
}

func Part1(f *os.File) {
	schematic := readSchematicFromFile(f)
	partNumberLocations := locateNumbers(&schematic)
	partNumberTotal := 0
	for _, location := range partNumberLocations {
		if schematic.isNumberAPart(location) {
			partNumberTotal += schematic.getNumberAtLocation(location)
		}
	}
	fmt.Printf("Sum of part numbers is %d\n", partNumberTotal)
}

func Part2(f *os.File) {
	schematic := readSchematicFromFile(f)
	partNumberLocations := locateNumbers(&schematic)
	partNumberMap := generatePartNumberMap(&schematic, &partNumberLocations)
	gearRatioTotal := 0
	for row := range schematic.contents {
		for col := range schematic.contents[row] {
			gearRatioTotal += getGearingRatio(&schematic, &partNumberMap, row, col)
		}
	}
	fmt.Printf("Sum of gearing ratios is %d\n", gearRatioTotal)
}
