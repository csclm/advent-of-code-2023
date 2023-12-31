package day23

import (
	"aoc-2023/aoc-lib"
	"fmt"
)

type HikingTrails struct {
	trailMap aoc.Grid[rune]
	start    aoc.Vec2
	end      aoc.Vec2
}

type HikingStep struct {
	location aoc.Vec2
	// lastStepDirection aoc.Vec2
}

func (ht HikingTrails) findLongestWalkLength() int {
	startStep := HikingStep{
		location: ht.start,
		// lastStepDirection: aoc.NewVec2(0, 1),
	}
	endStep := HikingStep{
		location: ht.end,
	}

	path := dijkstra(ht, startStep)
	dist := path.distances[endStep]

	pathGrid := ht.trailMap.Clone()
	currStep := endStep
	for {
		pathGrid.SetVec(currStep.location, 'O')
		ptr := path.links[currStep]
		if ptr == nil {
			break
		}
		currStep = *ptr
	}
	printableGrid := aoc.NewGrid[string](pathGrid.Width, pathGrid.Height)

	for loc := range pathGrid.Locations() {
		printableGrid.SetVec(loc, string(pathGrid.GetVec2(loc)))
	}

	printableGrid.Print()
	fmt.Println()
	return dist
}

// type HikingTrailGraph struct {
// 	vertices []HikingStep
// 	edges    map[int][]int
// }

// type HikingPath struct {
// 	location      aoc.Vec2
// 	lastDirection aoc.Vec2
// 	traversed     aoc.Set[aoc.Vec2]
// }

// func (ht HikingTrails) makeGraph() HikingTrailGraph {
// 	vertices := make([]HikingStep, 0)
// 	edges := make(map[int][]int)
// 	q := queue.New()
// 	q.Enqueue(HikingPath{
// 		location:  ht.start,
// 		lastDirection: aoc.NewVec2(0,1),
// 		traversed: aoc.NewSet[aoc.Vec2](),
// 	})
// 	for q.Len() != 0 {
// 		thisPath := q.Dequeue().(HikingPath)
// 		for _, dir := range aoc.CardinalDirections() {
// 			if isStepPossible(ht, thisPath.location, dir) {

// 			}
// 		}
// 	}
// }

func (ht HikingTrails) vertices() []HikingStep {
	result := make([]HikingStep, 0)
	for loc := range ht.trailMap.Locations() {
		if ht.trailMap.GetVec2(loc) == '#' {
			continue
		}
		result = append(result, HikingStep{
			location: loc,
			// lastStepDirection: dir,
		})
		// for _, dir := range aoc.CardinalDirections() {
		// 	if !isStepPossible(ht, loc.Plus(dir.Inverse()), dir) {
		// 		continue
		// 	}
		// 	result = append(result, HikingStep{
		// 		location: loc,
		// 		// lastStepDirection: dir,
		// 	})
		// }
	}
	return result
}

func (ht HikingTrails) neighbors(step HikingStep) []SearchEdge[HikingStep] {
	result := make([]SearchEdge[HikingStep], 0, 1)
	for _, dir := range aoc.CardinalDirections() {
		// if dir == step.lastStepDirection.Inverse() {
		// 	continue // can't backtrack
		// }
		if !isStepPossible(ht, step.location, dir) {
			continue
		}
		result = append(result, SearchEdge[HikingStep]{
			to: HikingStep{
				location: step.location.Plus(dir),
				// lastStepDirection: dir,
			},
			cost: -1, // negative cost to find the longest path
		})
	}
	return result
}

func isStepPossible(ht HikingTrails, location aoc.Vec2, direction aoc.Vec2) bool {
	sourceSquare, sourceInBounds := ht.trailMap.MaybeGetVec2(location)
	destSquare, destInBounds := ht.trailMap.MaybeGetVec2(location.Plus(direction))
	if !destInBounds || destSquare == '#' {
		return false
	}
	if !sourceInBounds || sourceSquare == '#' {
		return false
	}
	if sourceSquare != '.' && direction != directionFromArrowRune(sourceSquare) {
		return false
	}
	return true
}

func directionFromArrowRune(r rune) aoc.Vec2 {
	switch r {
	case '^':
		return aoc.NewVec2(0, -1)
	case 'v':
		return aoc.NewVec2(0, 1)
	case '<':
		return aoc.NewVec2(-1, 0)
	case '>':
		return aoc.NewVec2(1, 0)
	default:
		panic("invalid direction rune!")
	}
}
