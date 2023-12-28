package day18

import (
	"aoc-2023/aoc-lib"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mitchellh/iochan"
)

// Returns the area of the simple polygon defined as a list of ordered pairs
// https://en.wikipedia.org/wiki/Shoelace_formula
func shoelace(shape []aoc.Vec2) int {
	sum := 0
	for i := 0; i < len(shape); i++ {
		in := (i + 1) % len(shape)
		x := shape[i].X
		xn := shape[in].X
		y := shape[i].Y
		yn := shape[in].Y
		sum += x*yn - xn*y
	}
	return sum / 2
}

func orientationOfShape(shape []aoc.Vec2) int {
	if shoelace(shape) >= 0 {
		return 1
	} else {
		return -1
	}
}

// Return the direction in a right-handed coordinate system (up is the positive Y direction)
func directionVecFromRuneRightHanded(dirRune rune) aoc.Vec2 {
	switch dirRune {
	case 'U':
		return aoc.NewVec2(0, 1)
	case 'D':
		return aoc.NewVec2(0, -1)
	case 'R':
		return aoc.NewVec2(1, 0)
	case 'L':
		return aoc.NewVec2(-1, 0)
	default:
		panic("Invalid direction rune")
	}
}

func verticesFromDigInstructions(instructions []DigInstruction) []aoc.Vec2 {
	digPoints := make([]aoc.Vec2, 0, len(instructions))
	currentLocation := aoc.NewVec2(0, 0)
	for _, instruction := range instructions {
		digPoints = append(digPoints, currentLocation)
		dir := directionVecFromRuneRightHanded(instruction.direction)
		currentLocation = currentLocation.Plus(dir.Times(instruction.distance))
	}
	orientation := orientationOfShape(digPoints)
	result := make([]aoc.Vec2, 0, len(instructions))
	// Reread the instructions, accounting for the width of the trench (points along the exterior)
	currentLocation = aoc.NewVec2(0, 0)
	for i := range instructions {
		thisInstruction := instructions[i]
		nextInstruction := instructions[(i+1)%len(instructions)]
		entryDirection := directionVecFromRuneRightHanded(thisInstruction.direction)
		exitDirection := directionVecFromRuneRightHanded(nextInstruction.direction)
		cornerOffset := offsetForOrientation(entryDirection, exitDirection, orientation)

		thisDirection := directionVecFromRuneRightHanded(thisInstruction.direction)

		currentLocation = currentLocation.Plus(thisDirection.Times(thisInstruction.distance))
		result = append(result, currentLocation.Plus(cornerOffset))
	}
	return result
}

func offsetForOrientation(entryDirection, exitDirection aoc.Vec2, orientationOfShape int) aoc.Vec2 {
	inverseEntry := entryDirection.Inverse()
	// 1 if ccw, -1 if cw
	orientationOfTurn := inverseEntry.Y*exitDirection.X - inverseEntry.X*exitDirection.Y
	insideCornerOffset := inverseEntry.Plus(exitDirection)
	// Flip the offset to the other side if the curve runs in the direction of the shape
	offsetFromCenter := insideCornerOffset.Times(-1 * orientationOfTurn * orientationOfShape)
	// Relative to bottom-left corner
	return offsetFromCenter.Plus(aoc.NewVec2(1, 1)).DividedBy(2)
}

func parseInputWithHexInstructions(f *os.File) []DigInstruction {
	pattern := regexp.MustCompile(`([UDRL]) (\d+) \(#(\w+)\)`)
	result := make([]DigInstruction, 0)
	for line := range iochan.DelimReader(f, '\n') {
		matches := pattern.FindStringSubmatch(strings.TrimSpace(line))
		match := matches[3]
		distanceNum, _ := strconv.ParseInt(match[0:5], 16, 0)
		lastDigit := match[5]
		var direction rune
		switch lastDigit {
		case '0':
			direction = 'R'
		case '1':
			direction = 'D'
		case '2':
			direction = 'L'
		case '3':
			direction = 'U'
		default:
			panic("unknown direction digit")
		}
		result = append(result, DigInstruction{
			direction: direction,
			distance:  int(distanceNum),
			color:     "", // irrelevant
		})
	}
	return result
}
