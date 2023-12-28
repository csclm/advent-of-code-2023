package aoc

type Vec2 struct {
	X, Y int
}

func NewVec2(x, y int) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) Plus(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

func (v Vec2) Times(scalar int) Vec2 {
	return Vec2{v.X * scalar, v.Y * scalar}
}

func (v Vec2) DividedBy(scalar int) Vec2 {
	return Vec2{v.X / scalar, v.Y / scalar}
}

func (v Vec2) Inverse() Vec2 {
	return Vec2{-v.X, -v.Y}
}

func (v Vec2) IsInBoundingBox(width int, height int) bool {
	return v.X >= 0 && v.X < width && v.Y >= 0 && v.Y < height
}

func CardinalDirections() [4]Vec2 {
	return [4]Vec2{
		{0, 1},  // south
		{1, 0},  // east
		{0, -1}, // north
		{-1, 0}, // west
	}
}

func LocationsInGrid(width, height int) chan Vec2 {
	c := make(chan Vec2)
	go func() {
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				c <- NewVec2(col, row)
			}
		}
		close(c)
	}()
	return c
}
