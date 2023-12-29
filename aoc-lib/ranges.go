package aoc

type Range struct {
	Min int // inclusive
	Max int // exclusive
}

func (r Range) Plus(i int) Range {
	r.Max += i
	r.Min += i
	return r
}

func (r Range) Contains(input int) bool {
	return input >= r.Min && input < r.Max
}

func (r Range) Intersection(other Range) Range {
	if r.Min <= other.Min && r.Max >= other.Max {
		return other
	}
	if r.Min >= other.Min && r.Max <= other.Max {
		return r
	}
	if other.Contains(r.Min) {
		return Range{Min: r.Min, Max: other.Max}
	}
	if r.Contains(other.Min) {
		return Range{Min: other.Min, Max: r.Max}
	}
	return Range{Min: 0, Max: 0} // empty range
}

func (r Range) Size() int {
	return r.Max - r.Min
}

func (r Range) Intersects(other Range) bool {
	return r.Intersection(other).Size() > 0
}
