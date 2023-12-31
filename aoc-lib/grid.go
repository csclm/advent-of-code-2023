package aoc

import (
	"fmt"
	"slices"
)

type Grid[T any] struct {
	Width    int
	Height   int
	contents []T
}

func NewGrid[T any](width, height int) Grid[T] {
	return Grid[T]{
		Width:    width,
		Height:   height,
		contents: make([]T, width*height),
	}
}

func NewGridWithContents[T any](width, height int, contents []T) Grid[T] {
	return Grid[T]{
		Width:    width,
		Height:   height,
		contents: contents,
	}
}

func (g Grid[T]) Transposed() Grid[T] {
	result := NewGrid[T](g.Height, g.Width)
	for cell := range g.Locations() {
		result.Set(cell.Y, cell.X, g.GetVec2(cell))
	}
	return result
}

func (g Grid[T]) PrintRepr(repr func(T) rune) {
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			fmt.Print(repr(g.Get(c, r)))
		}
		fmt.Println()
	}
}

func (g Grid[rune]) Print() {
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			fmt.Print(g.Get(c, r))
		}
		fmt.Println()
	}
}

func (g Grid[T]) Locations() chan Vec2 {
	c := make(chan Vec2)
	go func() {
		for row := 0; row < g.Height; row++ {
			for col := 0; col < g.Width; col++ {
				c <- NewVec2(col, row)
			}
		}
		close(c)
	}()
	return c
}

func (g Grid[T]) Set(x, y int, value T) {
	g.SetVec(NewVec2(x, y), value)
}

func (g Grid[T]) SetVec(v Vec2, value T) {
	if !v.IsInBoundingBox(g.Width, g.Height) {
		panic("Location for set is out of bounds")
	}
	g.contents[v.Y*g.Width+v.X] = value
}

func (g Grid[T]) Get(x, y int) T {
	return g.GetVec2(NewVec2(x, y))
}

func (g Grid[T]) GetVec2(v Vec2) T {
	value, inBounds := g.MaybeGetVec2(v)
	if !inBounds {
		panic("Location was out of bounds for grid")
	}
	return value
}

func (g Grid[T]) MaybeGet(x, y int) (T, bool) {
	return g.MaybeGetVec2(NewVec2(x, y))
}

func (g Grid[T]) MaybeGetVec2(v Vec2) (T, bool) {
	if !v.IsInBoundingBox(g.Width, g.Height) {
		return *new(T), false
	}
	return g.contents[v.Y*g.Width+v.X], true
}

func (g Grid[T]) Clone() Grid[T] {
	g.contents = slices.Clone(g.contents)
	return g
}
