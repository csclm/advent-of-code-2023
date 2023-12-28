package day18

import (
	"aoc-2023/aoc-lib"
	"fmt"
	"slices"
)

type ShoelaceTestCase struct {
	name                    string
	instructions            []DigInstruction
	expectedVolume          int
	expectedExteriorCorners []aoc.Vec2
}

// Convex
/*
######
#....#
#....#
#....#
#....#
######
*/
var testCase1 ShoelaceTestCase = ShoelaceTestCase{
	name: "Simple Convex Square",
	instructions: []DigInstruction{
		{direction: 'U', distance: 5},
		{direction: 'R', distance: 5},
		{direction: 'D', distance: 5},
		{direction: 'L', distance: 5},
	},
	expectedVolume: 36,
	expectedExteriorCorners: []aoc.Vec2{
		aoc.NewVec2(0, 6),
		aoc.NewVec2(6, 6),
		aoc.NewVec2(6, 0),
		aoc.NewVec2(0, 0),
	},
}

// Concave
/*
###########
#.........#
#.........#
#......####
#......#...
#......#...
#......####
#.........#
#.........#
#.........#
###########
*/
var testCase2 ShoelaceTestCase = ShoelaceTestCase{
	name: "Simple Concave Shape",
	instructions: []DigInstruction{
		{direction: 'U', distance: 10},
		{direction: 'R', distance: 10},
		{direction: 'D', distance: 3},
		{direction: 'L', distance: 3},
		{direction: 'D', distance: 3},
		{direction: 'R', distance: 3},
		{direction: 'D', distance: 4},
		{direction: 'L', distance: 10},
	},
	expectedExteriorCorners: []aoc.Vec2{
		aoc.NewVec2(0, 11),
		aoc.NewVec2(11, 11),
		aoc.NewVec2(11, 7),
		aoc.NewVec2(8, 7),
		aoc.NewVec2(8, 5),
		aoc.NewVec2(11, 5),
		aoc.NewVec2(11, 0),
		aoc.NewVec2(0, 0),
	},
	expectedVolume: 115,
}

func runTestCases() {
	testCase(testCase1)
	testCase(testCase2)
}

func testCase(testCase ShoelaceTestCase) {
	vertices := verticesFromDigInstructions(testCase.instructions)
	area := aoc.IntAbs(shoelace(vertices))
	if !slices.Equal(vertices, testCase.expectedExteriorCorners) {
		fmt.Printf("Failed test case \"%s\" vertices\n", testCase.name)
		fmt.Print("expected: ")
		printVecSlice(testCase.expectedExteriorCorners)
		fmt.Println()
		fmt.Print("got.....: ")
		printVecSlice(vertices)
		fmt.Println()
	}
	if area != testCase.expectedVolume {
		fmt.Printf("Failed test case \"%s\" area\n", testCase.name)
		fmt.Printf("expected: %d", testCase.expectedVolume)
		fmt.Println()
		fmt.Printf("got.....: %d", area)
		fmt.Println()
	}
	fmt.Println()
}

func printVecSlice(vecSlice []aoc.Vec2) {
	fmt.Print("[")
	for _, vec := range vecSlice {
		fmt.Printf("(%d, %d), ", vec.X, vec.Y)
	}
	fmt.Print("]")
}
