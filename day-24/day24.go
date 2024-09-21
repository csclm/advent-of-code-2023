package day24

import (
	"fmt"
	"math"
	"os"
)

type FloatVec2 struct {
	X float64
	Y float64
}

type Hailstone struct {
	position FloatVec2
	velocity FloatVec2
}

type InequalityCondition struct {
}

func pathIntersectingAllHailstones() {
	// each hailstone has a position described by
	// p = dp*t + p0
	// call it a path
	// for 2 paths, what are all the lines that will intercept them?
	// ds*t + s0 = dp*t + p0
	// we need to make sure this has a solution for t > 0, so this imposes constraints on ds and s0
	// t = (p0 - s0) / (ds - dp)
	// ds > dp && s0 < p0 || ds < dp && s0 > p0
	// How do we generalize this to 3 dimensions?

}

func (h Hailstone) findPathIntersectionInFuture(other Hailstone) (FloatVec2, bool) {
	// TODO special case for vertical path?
	// Turned out not to be necessary for part 1

	m1 := h.velocity.Y / h.velocity.X
	m2 := other.velocity.Y / other.velocity.X
	b1 := h.position.Y - (h.position.X * m1)
	b2 := other.position.Y - (other.position.X * m2)
	x, y := linearIntercept(m1, b1, m2, b2)

	t1 := (x - h.position.X) / h.velocity.X
	t2 := (x - other.position.X) / other.velocity.X

	if t1 < 0 || t2 < 0 {
		return *new(FloatVec2), false
	}

	if math.IsInf(x, 0) {
		return *new(FloatVec2), false
	} else {
		return FloatVec2{x, y}, true
	}
}

// Intercept time, intercept value
func linearIntercept(m1, b1, m2, b2 float64) (float64, float64) {
	t := (b2 - b1) / (m1 - m2)
	val := m1*t + b1
	return t, val
}

func Part1(f *os.File) {
	stones := parseInput(f)
	intercepts := 0
	for i1, s1 := range stones {
		for i2 := 0; i2 < i1; i2++ {
			s2 := stones[i2]
			position, didIntercept := s1.findPathIntersectionInFuture(s2)
			if !didIntercept {
				continue
			}
			if position.X < 200000000000000 || position.Y < 200000000000000 || position.X > 400000000000000 || position.Y > 400000000000000 {
				continue
			}
			intercepts++
		}
	}

	fmt.Printf("Number of intercepts in range: %d", intercepts)
}
