package day7

import (
	"aoc-2023/aoc-lib"
	"os"
	"strings"
	"unicode"
)

type HandAndBid struct {
	hand Hand
	bid  int
}

func parseInput(f *os.File) []HandAndBid {
	result := make([]HandAndBid, 0)
	for line := range aoc.LineReader(f) {
		result = append(result, parseHandAndBid(line))
	}
	return result
}

func parseHandAndBid(line string) HandAndBid {
	components := strings.Split(line, " ")
	cards := [5]int{}
	for i, char := range components[0] {
		cards[i] = strengthOfRank(char)
	}
	bid := aoc.MustParseInt(components[1])
	return HandAndBid{
		hand: Hand{cards},
		bid:  bid,
	}
}

func strengthOfRank(rank rune) int {
	if unicode.IsDigit(rank) {
		return aoc.MustParseInt(string(rank))
	}
	switch rank {
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		panic("Invalid rank " + string(rank))
	}
}
