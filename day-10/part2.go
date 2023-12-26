package day10

import (
	"aoc-2023/aoc-lib"

	"github.com/golang-collections/collections/queue"
)

const MaskEmptySpace = PipeRune('.')
const MaskSolid = PipeRune('o')
const MaskFilled = PipeRune('X')

func CalculateEnclosedArea(grid *PipeGrid) int {
	filteredGrid := filterToLoop(grid)
	expandedGrid := makeExpandedMaskGrid(&filteredGrid)
	fillFromEdges(&expandedGrid)
	return countUnfilledExpandedCells(&expandedGrid)
}

// Starts at the edges and fills contiguous regions, replacing '.' with 'X'
func fillFromEdges(mask *PipeGrid) {
	for i := 0; i < mask.Width(); i++ {
		// no-op if already filled
		fillFromLocation(mask, aoc.NewVec2(i, 0))             // top edge
		fillFromLocation(mask, aoc.NewVec2(i, mask.Height())) // bottom edge
	}
	for i := 1; i < mask.Height()-1; i++ {
		// no-op if already filled
		fillFromLocation(mask, aoc.NewVec2(0, i))            // left edge
		fillFromLocation(mask, aoc.NewVec2(mask.Width(), i)) // right edge
	}
}

// For an expanded grid, counts the number of 3x3 grid cells without a pipe in them
func countUnfilledExpandedCells(expandedGrid *PipeGrid) int {
	emptyCells := 0
	for loc := range aoc.LocationsInGrid(expandedGrid.Width()/3, expandedGrid.Height()/3) {
		allAreEmpty := true
		for cellLoc := range aoc.LocationsInGrid(3, 3) {
			r, _ := expandedGrid.RuneAtLocation(loc.Times(3).Plus(cellLoc))
			if r != MaskEmptySpace {
				allAreEmpty = false
				break
			}
		}
		if allAreEmpty {
			emptyCells++
		}
	}
	return emptyCells
}

// For a grid with lots of pipes, deletes all the pipes that aren't connected to the main loop
func filterToLoop(grid *PipeGrid) PipeGrid {
	mask := PipeGridInit(grid.Width(), grid.Height(), MaskEmptySpace)
	for location := range LoopSteps(grid) {
		pipe, _ := grid.RuneAtLocation(location)
		mask.SetRuneAt(location, pipe)
	}
	return mask
}

// Expands the grid to allow filling to "slip between" pipes while keeping the loops closed
/*
Desired effect:
     . o .
L -> . o o
     . . .

     . . .
7 -> o o .
     . o .

     . . .
- -> o o o
     . . .

... etc
*/
func makeExpandedMaskGrid(loopGrid *PipeGrid) PipeGrid {
	expandedGrid := PipeGridInit(loopGrid.Width()*3, loopGrid.Height()*3, MaskEmptySpace)
	for loc := range aoc.LocationsInGrid(loopGrid.Width(), loopGrid.Height()) {
		thisPipe, _ := loopGrid.RuneAtLocation(loc)
		if thisPipe == '.' {
			continue
		}
		connections := thisPipe.Connections()
		expandedGrid.SetRuneAt(loc.Times(3).Plus(aoc.NewVec2(1, 1)), MaskSolid)
		if connections.north {
			expandedGrid.SetRuneAt(loc.Times(3).Plus(aoc.NewVec2(1, 0)), MaskSolid)
		}
		if connections.south {
			expandedGrid.SetRuneAt(loc.Times(3).Plus(aoc.NewVec2(1, 2)), MaskSolid)
		}
		if connections.east {
			expandedGrid.SetRuneAt(loc.Times(3).Plus(aoc.NewVec2(2, 1)), MaskSolid)
		}
		if connections.west {
			expandedGrid.SetRuneAt(loc.Times(3).Plus(aoc.NewVec2(0, 1)), MaskSolid)
		}
	}
	return expandedGrid
}

// Takes a PipeGrid mask and fills in a contiguous region that encloses the start point
func fillFromLocation(mask *PipeGrid, start aoc.Vec2) {
	startRune, startInBounds := mask.RuneAtLocation(start)
	if !startInBounds {
		return
	}
	if startRune != MaskEmptySpace {
		return
	}
	mask.SetRuneAt(start, MaskFilled)
	q := queue.New()
	q.Enqueue(start)
	for q.Len() != 0 {
		location := q.Dequeue().(aoc.Vec2)
		for _, dir := range aoc.CardinalDirections() {
			nextLocation := location.Plus(dir)
			nextRune, nextIsInBounds := mask.RuneAtLocation(nextLocation)
			if nextIsInBounds && nextRune == MaskEmptySpace {
				q.Enqueue(nextLocation)
				mask.SetRuneAt(nextLocation, MaskFilled)
			}
		}
	}
}
