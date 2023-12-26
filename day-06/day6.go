package day6

import (
	"fmt"
	"math"
)

type Race struct {
	time     int
	distance int
}

// Skipping the parsing logic on this one
// Ordering turns out to be important for part 2
var Races [4]Race = [4]Race{
	{time: 54, distance: 446},
	{time: 81, distance: 1292},
	{time: 70, distance: 1035},
	{time: 88, distance: 1007},
}

func Part1() {
	answer := 1
	for _, race := range Races {
		answer *= race.numberOfWaysToWin()
	}
	fmt.Printf("Part 1 answer is %d\n", answer)
}

func Part2() {
	longRace := Race{}
	for _, race := range Races {
		longRace.time = rightConcat(longRace.time, race.time)
		longRace.distance = rightConcat(longRace.distance, race.distance)
	}
	fmt.Printf("Part 2 answer is %d\n", longRace.numberOfWaysToWin())
}

func rightConcat(left int, right int) int {
	shift := int(math.Pow10(1 + int(math.Log10(float64(right)))))
	return left*shift + right
}

func (r Race) numberOfWaysToWin() int {
	// For X ms held and Y distance, what is the time
	// time = held + dist/held
	// time relative to record = held + dist/held - record = 0
	// held + dist/held - record = 0
	// multiply by held (h) h^2 - rh + d = 0
	// quadratic formula time! a = 1 b = -r c = dist
	x1, x2 := quadratic(1, -float64(r.time), float64(r.distance))
	low := int(math.Ceil(min(x1, x2)))
	high := int(math.Floor(max(x1, x2)))
	return high - low + 1
}

func quadratic(a float64, b float64, c float64) (float64, float64) {
	x1 := (-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	x2 := (-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	return x1, x2
}
