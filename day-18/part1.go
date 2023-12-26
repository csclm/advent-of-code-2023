package day18

import (
	"aoc-2023/aoc-lib"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/queue"
	"github.com/mitchellh/iochan"
)

type GroundCell struct {
	dug   bool
	color string
}

type Hole struct {
	location aoc.Vec2
	color    string
}

type DigInstruction struct {
	direction rune
	distance  int
	color     string // hexadecimal
}

func digInteriorHoles(ground [][]GroundCell) {
	width := len(ground[0])
	height := len(ground)

	mask := make([][]bool, height)
	for i := 0; i < height; i++ {
		mask[i] = make([]bool, width)
	}

	q := queue.New()
	for _, coord := range perimiterCoords(width, height) {
		if ground[coord.Y][coord.X].dug {
			continue
		}
		for _, dir := range aoc.CardinalDirections() {
			newCoord := coord.Plus(dir)
			if newCoord.IsInBoundingBox(width, height) &&
				!mask[newCoord.Y][newCoord.X] &&
				!ground[newCoord.Y][newCoord.X].dug {
				q.Enqueue(newCoord)
			}
		}
	}
	for q.Len() != 0 {
		coord := q.Dequeue().(aoc.Vec2)
		for _, dir := range aoc.CardinalDirections() {
			newCoord := coord.Plus(dir)
			if newCoord.IsInBoundingBox(width, height) &&
				!mask[newCoord.Y][newCoord.X] &&
				!ground[newCoord.Y][newCoord.X].dug {
				mask[newCoord.Y][newCoord.X] = true
				q.Enqueue(newCoord)
			}
		}
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if !mask[row][col] {
				ground[row][col].dug = true
			}
		}
	}
}

func digHoles(instructions []DigInstruction) []Hole {
	result := make([]Hole, 0, len(instructions))
	currentLocation := aoc.NewVec2(0, 0)

	// TODO what color should the initial hole be?
	result = append(result, Hole{location: currentLocation, color: ""})
	for _, instruction := range instructions {
		var digDirection aoc.Vec2
		switch instruction.direction {
		case 'U':
			digDirection = aoc.NewVec2(0, -1)
		case 'D':
			digDirection = aoc.NewVec2(0, 1)
		case 'R':
			digDirection = aoc.NewVec2(1, 0)
		case 'L':
			digDirection = aoc.NewVec2(-1, 0)
		}
		for i := 0; i < instruction.distance; i++ {
			currentLocation = currentLocation.Plus(digDirection)
			result = append(result, Hole{location: currentLocation, color: instruction.color})
		}
	}
	return result
}

func makeGridFromHoles(holes []Hole) [][]GroundCell {
	minX, minY, maxX, maxY := 999999, 999999, -999999, -999999
	for _, hole := range holes {
		minX = min(minX, hole.location.X)
		minY = min(minY, hole.location.Y)
		maxX = max(maxX, hole.location.X)
		maxY = max(maxY, hole.location.Y)
	}
	width := maxX - minX + 1
	height := maxY - minY + 1
	result := make([][]GroundCell, height)
	for i := 0; i < height; i++ {
		result[i] = make([]GroundCell, width)
	}
	for _, hole := range holes {
		result[hole.location.Y-minY][hole.location.X-minX] = GroundCell{
			dug:   true,
			color: hole.color,
		}
	}
	return result
}

func parseInput(f *os.File) []DigInstruction {
	pattern := regexp.MustCompile(`([UDRL]) (\d+) \((#\w+)\)`)
	result := make([]DigInstruction, 0)
	for line := range iochan.DelimReader(f, '\n') {
		match := pattern.FindStringSubmatch(strings.TrimSpace(line))
		distanceNum, _ := strconv.ParseInt(match[2], 10, 0)
		result = append(result, DigInstruction{
			direction: []rune(match[1])[0],
			distance:  int(distanceNum),
			color:     match[3],
		})
	}
	return result
}
