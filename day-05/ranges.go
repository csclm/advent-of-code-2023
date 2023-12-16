package main

type Range struct {
	min int // inclusive
	max int // exclusive
}

func (r Range) Contains(input int) bool {
	return input >= r.min && input < r.max
}

func (r Range) Intersection(other Range) (Range, bool) {
	if r.min <= other.min && r.max >= other.max {
		return other, true
	}
	if r.min >= other.min && r.max <= other.max {
		return r, true
	}
	if other.Contains(r.min) {
		return Range{min: r.min, max: other.max}, true
	}
	if r.Contains(other.min) {
		return Range{min: other.min, max: r.max}, true
	}
	return Range{min: 0, max: 0}, false
}
