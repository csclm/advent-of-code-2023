package aoc

type Range struct {
	Min int // inclusive
	Max int // exclusive
}

func (r Range) Contains(input int) bool {
	return input >= r.Min && input < r.Max
}

func (r Range) Intersection(other Range) (Range, bool) {
	if r.Min <= other.Min && r.Max >= other.Max {
		return other, true
	}
	if r.Min >= other.Min && r.Max <= other.Max {
		return r, true
	}
	if other.Contains(r.Min) {
		return Range{Min: r.Min, Max: other.Max}, true
	}
	if r.Contains(other.Min) {
		return Range{Min: other.Min, Max: r.Max}, true
	}
	return Range{Min: 0, Max: 0}, false
}
