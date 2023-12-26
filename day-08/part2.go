package day8

import (
	"aoc-2023/aoc-lib"
	"slices"

	"github.com/golang-collections/collections/set"
)

type PathMap struct {
	directions string
	network    Network
}

type PathPointer struct {
	directionStep int
	node          string
}

func (p PathMap) Next(ptr PathPointer) PathPointer {
	if p.directions[ptr.directionStep] == 'L' {
		return PathPointer{(ptr.directionStep + 1) % len(p.directions), p.network.nodes[ptr.node].left}
	} else {
		return PathPointer{(ptr.directionStep + 1) % len(p.directions), p.network.nodes[ptr.node].right}
	}
}

/*
Represents the path that a visitor can take through a directional graph
Every path starts with some nodes visited followed by a cycle
for instance, 1 2 3 4 5 6 4 5 6 4 5 6 ...
Let's represent this as 1 2 3 [4 5 6]

We can annotate which nodes in this cycle are destinations
If odd nodes are our destinations in the example above, this would look like
cycleStart: 3
cycleLength: 3
visitsBeforeCycle: [ 0, 2 ]
visitsInCycle: [ 1 ]
*/
type VisitPath struct {
	cycleStart        int
	cycleLength       int
	visitsBeforeCycle []int
	visitsInCycle     []int
}

func calculatePathFromStartPoint(start string, directions string, network Network) VisitPath {
	pathMap := PathMap{directions, network}
	stepsAtEndNode := make([]int, 0)

	// Figure out the presence of a cycle, marking finishing nodes along the way
	slowPointer := PathPointer{0, start}
	slowPointerStep := 0
	fastPointer := PathPointer{0, start}
	for slowPointerStep == 0 || slowPointer != fastPointer {
		if isFinishingNode(slowPointer.node) {
			stepsAtEndNode = append(stepsAtEndNode, slowPointerStep)
		}
		slowPointerStep++
		slowPointer = pathMap.Next(slowPointer)
		fastPointer = pathMap.Next(pathMap.Next(fastPointer))
	}

	cycleLength := slowPointerStep

	// Find the start of the cycle
	fastPointer = PathPointer{0, start}
	cycleStartStep := 0
	for slowPointer != fastPointer {
		if isFinishingNode(slowPointer.node) {
			stepsAtEndNode = append(stepsAtEndNode, slowPointerStep)
		}
		fastPointer = pathMap.Next(fastPointer)
		slowPointer = pathMap.Next(slowPointer)
		slowPointerStep++
		cycleStartStep++
	}

	// Split visits into "before-cycle" vs "in-cycle"
	endNodeVisitsInCycle := make([]int, 0)
	endNodeVisitsBeforeCycle := make([]int, 0)
	for _, endNodeStep := range stepsAtEndNode {
		stepInCycle := endNodeStep - cycleStartStep
		if stepInCycle >= 0 {
			endNodeVisitsInCycle = append(endNodeVisitsInCycle, stepInCycle)
		} else {
			endNodeVisitsBeforeCycle = append(endNodeVisitsBeforeCycle, endNodeStep)
		}
	}

	return VisitPath{
		cycleStart:        cycleStartStep,
		cycleLength:       cycleLength,
		visitsInCycle:     endNodeVisitsInCycle,
		visitsBeforeCycle: endNodeVisitsBeforeCycle,
	}
}

func isStartingNode(nodeName string) bool {
	return nodeName[2] == 'A'
}

func isFinishingNode(nodeName string) bool {
	return nodeName[2] == 'Z'
}

// For 2 visit paths, generates a virtual path representing the points where
// destination nodes were visited by both paths on the same step
func (path VisitPath) Join(other VisitPath) VisitPath {
	var early VisitPath // Starts its cycle early
	var late VisitPath  // Starts its cycle at the same time or later as "early"
	if path.cycleStart < other.cycleStart {
		early = path
		late = other
	} else {
		early = other
		late = path
	}
	visitsBeforeCycle := make([]int, 0)
	for _, step := range late.visitsBeforeCycle {
		if step < early.cycleStart {
			if slices.Contains(early.visitsBeforeCycle, step) {
				visitsBeforeCycle = append(visitsBeforeCycle, step)
			}
		} else {
			if slices.Contains(early.visitsInCycle, (step-early.cycleStart)%early.cycleLength) {
				visitsBeforeCycle = append(visitsBeforeCycle, step)
			}
		}
	}

	cycleLength := aoc.Lcm(early.cycleLength, late.cycleLength)
	earlyMultiple := cycleLength / early.cycleLength
	lateMultiple := cycleLength / late.cycleLength

	visitsInCycle := make([]int, 0)
	visitsFromEarlyCycle := set.New()
	for i := 1; i <= lateMultiple; i++ {
		for _, step := range late.visitsInCycle {
			visitsFromEarlyCycle.Insert(step)
		}
	}
	for i := 1; i <= earlyMultiple; i++ {
		for _, step := range early.visitsInCycle {
			offset := late.cycleStart - early.cycleStart
			stepInCycle := (step*i - offset) % cycleLength
			if stepInCycle < 0 {
				stepInCycle += cycleLength // Force it into [0,cycleLength)
			}
			if visitsFromEarlyCycle.Has(stepInCycle) {
				visitsInCycle = append(visitsInCycle, stepInCycle)
			}
		}
	}

	return VisitPath{
		cycleStart:        late.cycleStart,
		cycleLength:       cycleLength,
		visitsBeforeCycle: visitsBeforeCycle,
		visitsInCycle:     visitsInCycle,
	}

}

func identityPath() VisitPath {
	return VisitPath{
		cycleStart:        0,
		cycleLength:       1,
		visitsBeforeCycle: []int{},
		visitsInCycle:     []int{0},
	}
}

func numberOfTraversalStepsPart2(directions string, network Network) int {
	truncatedDirections := truncateSymmetricalString(directions)
	paths := make([]VisitPath, 0)
	for k := range network.nodes {
		if isStartingNode(k) {
			thisNodePath := calculatePathFromStartPoint(k, truncatedDirections, network)
			paths = append(paths, thisNodePath)
		}
	}

	path := identityPath()
	for _, individualPath := range paths {
		path = path.Join(individualPath)
	}

	if len(path.visitsBeforeCycle) != 0 {
		minStep := 99999
		for _, step := range path.visitsBeforeCycle {
			minStep = min(minStep, step)
		}
		return minStep
	} else if len(path.visitsInCycle) != 0 {
		minStep := 99999
		for _, step := range path.visitsInCycle {
			minStep = min(minStep, step)
		}
		return minStep + path.cycleStart
	} else {
		panic("solution path was empty!")
	}
}
