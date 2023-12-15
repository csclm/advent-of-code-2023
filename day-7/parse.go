package main

import (
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/mitchellh/iochan"
)

type HandAndBid struct {
	hand Hand
	bid  int
}

func parseInput(f *os.File) []HandAndBid {
	result := make([]HandAndBid, 0)
	for line := range iochan.DelimReader(f, '\n') {
		result = append(result, parseHandAndBid(line))
	}
	return result
}

func parseHandAndBid(line string) HandAndBid {
	components := strings.Split(strings.TrimSpace(line), " ")
	cards := [5]int{}
	for i, char := range components[0] {
		cards[i] = strengthOfRank(char)
	}
	bid := mustParseInt(components[1])
	return HandAndBid{
		hand: Hand{cards},
		bid:  bid,
	}
}

func mustParseInt(s string) int {
	num, _ := strconv.ParseInt(s, 10, 0)
	return int(num)
}

func strengthOfRank(rank rune) int {
	if unicode.IsDigit(rank) {
		return mustParseInt(string(rank))
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
