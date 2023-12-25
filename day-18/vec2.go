package main

type Vec2 struct {
	x, y int
}

func (v Vec2) Plus(other Vec2) Vec2 {
	return Vec2{v.x + other.x, v.y + other.y}
}

func (v Vec2) IsInBounds(width, height int) bool {
	return v.x >= 0 && v.y >= 0 && v.x < width && v.y < height
}

func perimiterCoords(width, height int) []Vec2 {
	result := make([]Vec2, 0)
	for i := 0; i < height; i++ {
		result = append(result, Vec2{0, i})
		result = append(result, Vec2{width - 1, i})
	}
	for i := 1; i < width-1; i++ {
		result = append(result, Vec2{i, 0})
		result = append(result, Vec2{i, height - 1})
	}
	return result
}

func cardinalDirections() []Vec2 {
	return []Vec2{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
}
