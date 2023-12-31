package main

import (
	day1 "aoc-2023/day-01"
	day2 "aoc-2023/day-02"
	day3 "aoc-2023/day-03"
	day4 "aoc-2023/day-04"
	day5 "aoc-2023/day-05"
	day6 "aoc-2023/day-06"
	day7 "aoc-2023/day-07"
	day8 "aoc-2023/day-08"
	day9 "aoc-2023/day-09"
	day10 "aoc-2023/day-10"
	day11 "aoc-2023/day-11"
	day12 "aoc-2023/day-12"
	day13 "aoc-2023/day-13"
	day14 "aoc-2023/day-14"
	day15 "aoc-2023/day-15"
	day16 "aoc-2023/day-16"
	day17 "aoc-2023/day-17"
	day18 "aoc-2023/day-18"
	day19 "aoc-2023/day-19"
	day20 "aoc-2023/day-20"
	day21 "aoc-2023/day-21"
	day24 "aoc-2023/day-24"

	// day23 "aoc-2023/day-23"
	// day25 "aoc-2023/day-25"
	"os"
	"strconv"
)

func main() {
	CurrentPart()
	// AllParts()
}

func CurrentPart() {
	day24.Part1(input(24))
}

func AllParts() {
	day1.Part1(input(1))
	day1.Part2(input(1))

	day2.Part1(input(2))
	day2.Part2(input(2))

	day3.Part1(input(3))
	day3.Part2(input(3))

	day4.Part1(input(4))
	day4.Part2(input(4))

	day5.Part1(input(5))
	day5.Part2(input(5))

	day6.Part1(input(6))
	day6.Part2(input(6))

	day7.Part1(input(7))
	day7.Part2(input(7))

	day8.Part1(input(8))
	//day8.Part2(input(8))

	day9.Part1(input(9))
	day9.Part2(input(9))

	day10.Part1(input(10))
	day10.Part2(input(10))

	day11.Part1(input(11))
	day11.Part2(input(11))

	day12.Part1(input(12))
	day12.Part2(input(12))

	day13.Part1(input(13))
	day13.Part2(input(13))

	day14.Part1(input(14))
	//day14.Part2(input(1))

	day15.Part1(input(15))
	day15.Part2(input(15))

	day16.Part1(input(16))
	day16.Part2(input(16))

	day17.Part1(input(17))
	day17.Part2(input(17))

	day18.Part1(input(18))
	//day18.Part2(input(1))

	day19.Part1(input(19))
	// day19.Part2(input(19))

	day20.Part1(input(20))
	// day20.Part2(input(20))

	day21.Part1(input(21))
	// day21.Part2(input(21))

	// day22.Part1(input(22))
	// day22.Part2(input(22))

	// day23.Part1(input(23))
	// day23.Part2(input(23))

	day24.Part1(input(24))
	// day24.Part2(input(24))

	// day25.Part1(input(25))
	// day25.Part2(input(25))
}

func input(day int) *os.File {
	str := strconv.FormatInt(int64(day), 10)
	if len(str) == 1 {
		str = "0" + str
	}
	f, _ := os.Open("./inputs/day-" + str + ".txt")
	return f
}
