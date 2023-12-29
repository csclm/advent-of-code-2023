package aoc

import "maps"

// Typed set - more ergonomic than the untyped set.Set

type Set[T comparable] struct {
	contents map[T]byte
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		contents: make(map[T]byte),
	}
}

func (s Set[T]) Insert(element T) {
	s.contents[element] = 1
}

func (s Set[T]) InsertAll(other Set[T]) {
	for e := range other.Elements() {
		s.Insert(e)
	}
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := NewSet[T]()
	result.InsertAll(s)
	result.InsertAll(other)
	return result
}

func (s Set[T]) TakeOne() T {
	for k := range s.contents {
		return k
	}
	panic("attempt to take from empty set")
}

func (s Set[T]) Delete(element T) {
	delete(s.contents, element)
}

func (s Set[T]) Has(element T) bool {
	_, has := s.contents[element]
	return has
}

func (s Set[T]) Elements() chan T {
	c := make(chan T)
	go func() {
		for k := range s.contents {
			c <- k
		}
		close(c)
	}()
	return c
}

func (s Set[T]) AsSlice() []T {
	result := make([]T, 0)
	for k := range s.contents {
		result = append(result, k)
	}
	return result
}

func (s Set[T]) Clone() Set[T] {
	s.contents = maps.Clone(s.contents)
	return s
}

func (s Set[T]) Len() int {
	return len(s.contents)
}
