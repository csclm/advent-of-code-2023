package main

type Vec2 struct {
	x, y int
}

func (v Vec2) Plus(other Vec2) Vec2 {
	return Vec2{v.x + other.x, v.y + other.y}
}

func cardinalDirections() [4]Vec2 {
	return [4]Vec2{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
}
