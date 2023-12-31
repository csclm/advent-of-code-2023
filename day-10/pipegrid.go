package day10

import "aoc-2023/aoc-lib"

type PipeGrid [][]PipeRune

type PipeRune rune

type PipeConnections struct {
	north, south, east, west bool
}

func (grid PipeGrid) Width() int {
	return len(grid[0])
}

func (grid PipeGrid) Height() int {
	return len(grid)
}

func (c PipeConnections) PipeRune() PipeRune {
	if c.north && c.south {
		return '|'
	}
	if c.east && c.west {
		return '-'
	}
	if c.north && c.east {
		return 'L'
	}
	if c.north && c.west {
		return 'J'
	}
	if c.south && c.east {
		return 'F'
	}
	if c.south && c.west {
		return '7'
	}
	panic("Pipe rune not defined for given connections")
}

func (pr PipeRune) Connections() PipeConnections {
	switch pr {
	case 'S':
		return PipeConnections{true, true, true, true} // could connect to any direction
	case 'L':
		return PipeConnections{true, false, true, false}
	case '7':
		return PipeConnections{false, true, false, true}
	case '|':
		return PipeConnections{true, true, false, false}
	case '-':
		return PipeConnections{false, false, true, true}
	case 'F':
		return PipeConnections{false, true, true, false}
	case 'J':
		return PipeConnections{true, false, false, true}
	default:
		panic("Invalid piperune")
	}
}

func (pr PipeRune) ConnectsTo(direction aoc.Vec2) bool {
	connections := pr.Connections()
	if direction == (aoc.NewVec2(1, 0)) {
		return connections.east
	} else if direction == (aoc.NewVec2(-1, 0)) {
		return connections.west
	} else if direction == (aoc.NewVec2(0, 1)) {
		return connections.south
	} else if direction == (aoc.NewVec2(0, -1)) {
		return connections.north
	} else {
		return false
	}
}

func (grid PipeGrid) MustFindStartingPoint() aoc.Vec2 {
	for loc := range aoc.LocationsInGrid(grid.Width(), grid.Height()) {
		r, _ := grid.RuneAtLocation(loc)
		if r == 'S' {
			return loc
		}
	}
	panic("PipeGrid didn't find starting point!")
}

// rune, in bounds?
func (grid PipeGrid) RuneAtLocation(location aoc.Vec2) (PipeRune, bool) {
	if location.X >= 0 && location.X < grid.Width() && location.Y >= 0 && location.Y < grid.Height() {
		return PipeRune(grid[location.Y][location.X]), true
	}
	return '\x00', false
}

func (grid PipeGrid) SetRuneAt(location aoc.Vec2, r PipeRune) {
	if location.X >= 0 && location.X < grid.Width() && location.Y >= 0 && location.Y < grid.Height() {
		grid[location.Y][location.X] = r
	}
}

func PipeGridInit(width, height int, r PipeRune) PipeGrid {
	arr := make([][]PipeRune, height)
	for row := 0; row < height; row++ {
		arr[row] = make([]PipeRune, width)
		for col := 0; col < width; col++ {
			arr[row][col] = r
		}
	}
	return PipeGrid(arr)
}
