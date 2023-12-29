package day21

import "aoc-2023/aoc-lib"

/*
TODO figure out how to do a sort of dynamic programming problem with this
Regular dynamic programming isn't going to work well because it has to list out all the locations
of which there might be billions

I'd have to find a way to collapse subproblems to take advantage of the repeating grid

mainly I have to worry about steps being too big - 26501365 is probably going to overflow the stack
I need to figure out how to collapse this number down.. or maybe I just iterate instead of recurse?


I wonder if it involves computing reachable locations in X steps for every square in the grid
*/
func reachableLocations(grid [][]rune, steps int, start aoc.Vec2) {

}

type StepsSubproblem struct {
	steps int
	start aoc.Vec2
}

/*
Have to break this down into solvable subproblems

maybe the first computation could be - for any starting position within the map, how many steps does it take to reach
any square on any adjacent instance of the map?

*/

// func reachableLocationsMemo(grid [][]rune, steps int, start Vec2, memo map[StepsSubproblem]([]Vec2)) []Vec2 {
// 	if steps == 0 {
// 		return []Vec2{start}
// 	}

// 	result := make([]Vec2, 0)
// 	for _, dir := range cardinalDirections() {
// 		location := start.Plus(dir)
// 		if grid[location]
// 	}

// }
