package day1

import (
	"aoc-2023/aoc-lib"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func Part1(f *os.File) {
	total1 := 0
	for line := range aoc.LineReader(f) {
		total1 += numFromLinePart1(line)
	}
	fmt.Printf("Total 1 is: %d\n", total1)
}

func Part2(f *os.File) {
	total2 := 0
	for line := range aoc.LineReader(f) {
		total2 += numFromLinePart2(line)
	}
	fmt.Printf("Total 2 is: %d\n", total2)
}

func numFromLinePart1(line string) int {
	firstRune := ' '
	lastRune := ' '
	for _, char := range line {
		if !unicode.IsNumber(char) {
			continue
		}
		if firstRune == ' ' {
			firstRune = char
		}
		lastRune = char
	}
	if firstRune == ' ' {
		return 0
	}
	firstDigit, _ := strconv.ParseInt(string(firstRune), 10, 8)
	lastDigit, _ := strconv.ParseInt(string(lastRune), 10, 8)
	return int(firstDigit)*10 + int(lastDigit)
}

func numFromLinePart2(line string) int {
	// Normally I would use a positive lookahead group to allow for overlapping matches
	// E.g. (?=([0-9]|one|two|three|four|five|six|seven|eight|nine))

	// but golang doesn't support it, so I'm taking advantange of the fact that
	// the regex engine finds the leftmost non-overlapping match, and just matching on
	// the reverse of the string to find the rightmost non-overlapping match
	forwardPattern, _ := regexp.Compile(`([0-9]|one|two|three|four|five|six|seven|eight|nine)`)
	reversePattern, _ := regexp.Compile(`([0-9]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`)

	forwardMatches := forwardPattern.FindAllString(line, -1)
	reverseMatches := reversePattern.FindAllString(aoc.ReverseString(line), -1)

	firstDigit := digitFromMatch(forwardMatches[0])
	lastDigit := digitFromMatch(aoc.ReverseString(reverseMatches[0]))

	if firstDigit == -1 || lastDigit == -1 {
		panic("shouldn't happen")
	}

	return firstDigit*10 + lastDigit
}

func digitFromMatch(match string) int {
	if unicode.IsNumber(rune(match[0])) {
		digit, _ := strconv.ParseInt(match[0:1], 10, 8)
		return int(digit)
	}
	switch match {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}
	return -1
}
