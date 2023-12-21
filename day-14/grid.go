package main

const (
	North = iota
	East  = iota
	South = iota
	West  = iota
)

type Grid struct {
	contents    [][]rune
	orientation int // represents how the contents are oriented
}

func (g Grid) RotatedRight() Grid {
	var newDirection int
	switch g.orientation {
	case North:
		newDirection = East
	case South:
		newDirection = West
	case East:
		newDirection = South
	case West:
		newDirection = North
	default:
		panic("Invalid direction")
	}
	g.orientation = newDirection
	return g
}

func (g Grid) RotatedLeft() Grid {
	var newDirection int
	switch g.orientation {
	case North:
		newDirection = West
	case South:
		newDirection = East
	case East:
		newDirection = North
	case West:
		newDirection = South
	default:
		panic("Invalid direction")
	}
	g.orientation = newDirection
	return g
}

func (g Grid) RuneAt(row int, col int) rune {
	width := len(g.contents[0])
	height := len(g.contents)
	switch g.orientation {
	case North:
		return g.contents[row][col]
	case East:
		return g.contents[col][width-row-1]
	case South:
		return g.contents[height-row-1][width-col-1]
	case West:
		return g.contents[height-col-1][row]
	default:
		panic("Invalid orientation")
	}
}

func (g Grid) SetRuneAt(row int, col int, r rune) {
	width := len(g.contents[0])
	height := len(g.contents)
	switch g.orientation {
	case North:
		g.contents[row][col] = r
	case East:
		g.contents[col][width-row-1] = r
	case South:
		g.contents[height-row-1][width-col-1] = r
	case West:
		g.contents[height-col-1][row] = r
	default:
		panic("Invalid orientation")
	}
}

func (g Grid) Width() int {
	switch g.orientation {
	case North:
		fallthrough
	case South:
		return len(g.contents[0])
	case East:
		fallthrough
	case West:
		return len(g.contents)
	default:
		panic("Invalid orientation")
	}
}

func (g Grid) Height() int {
	switch g.orientation {
	case North:
		fallthrough
	case South:
		return len(g.contents)
	case East:
		fallthrough
	case West:
		return len(g.contents[0])
	default:
		panic("Invalid orientation")
	}
}
