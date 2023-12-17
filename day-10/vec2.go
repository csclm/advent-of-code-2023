package main

type Vec2 struct {
	x int
	y int
}

func (v Vec2) Plus(other Vec2) Vec2 {
	return Vec2{v.x + other.x, v.y + other.y}
}

func (v Vec2) Times(scalar int) Vec2 {
	return Vec2{v.x * scalar, v.y * scalar}
}

func (v Vec2) Inverse() Vec2 {
	return Vec2{-v.x, -v.y}
}

func cardinalDirections() []Vec2 {
	return []Vec2{
		{0, 1},  // south
		{1, 0},  // east
		{0, -1}, // north
		{-1, 0}, // west
	}
}

func locationsInGrid(width int, height int) chan Vec2 {
	c := make(chan Vec2)
	go func() {
		for row := 0; row < width; row++ {
			for col := 0; col < height; col++ {
				c <- Vec2{col, row}
			}
		}
		close(c)
	}()
	return c
}
