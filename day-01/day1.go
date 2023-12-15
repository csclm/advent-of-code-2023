package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"

	"github.com/mitchellh/iochan"
)

func main() {
	f, _ := os.Open("./input.txt")
	total1 := 0
	//total2 := 0
	for line := range iochan.DelimReader(f, '\n') {
		total1 += numFromLinePart1(line)
		//total2 += numFromLinePart2(line)
	}
	fmt.Printf("Total 1 is: %d\n", total1)
	//fmt.Printf("Total 2 is: %d\n", total2)
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
	// This would work if golang supported ?=
	pattern, _ := regexp.Compile("(?=([0-9]|one|two|three|four|five|six|seven|eight|nine))")

	firstDigit := -1
	lastDigit := -1

	matches := pattern.FindAllString(line, -1)

	for _, match := range matches {
		theDigit := digitFromMatch(match)
		if theDigit == -1 {
			panic("shouldn't happen")
		}
		if firstDigit == -1 {
			firstDigit = theDigit
		}
		lastDigit = theDigit
	}

	if firstDigit == -1 {
		return 0
	}

	return int(firstDigit)*10 + int(lastDigit)
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
