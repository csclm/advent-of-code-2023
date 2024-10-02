package day4

import (
	"aoc-2023/aoc-lib"
	"os"
	"regexp"
)

func parseInput(f *os.File) []Card {
	cards := make([]Card, 0)
	for line := range aoc.LineReader(f) {
		card := parseCard(line)
		cards = append(cards, card)
	}
	return cards
}

func parseCard(line string) Card {
	pattern := regexp.MustCompile(`\d+`)
	result := Card{}
	for i, match := range pattern.FindAllString(line, -1) {
		num := aoc.MustParseInt(match)
		if i == 0 {
			result.cardNum = num
		} else if i < len(result.winningNumbers)+1 {
			result.winningNumbers[i-1] = num
		} else {
			result.numbersOnCard[i-1-len(result.winningNumbers)] = num
		}
	}
	return result
}
