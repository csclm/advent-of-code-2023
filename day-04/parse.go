package day4

import (
	"os"
	"regexp"
	"strconv"

	"github.com/mitchellh/iochan"
)

func parseInput(f *os.File) []Card {
	cards := make([]Card, 0)
	for line := range iochan.DelimReader(f, '\n') {
		card := parseCard(line)
		cards = append(cards, card)
	}
	return cards
}

func parseCard(line string) Card {
	pattern := regexp.MustCompile(`\d+`)
	result := Card{}
	for i, match := range pattern.FindAllString(line, -1) {
		num, _ := strconv.ParseInt(match, 10, 0)
		if i == 0 {
			result.cardNum = int(num)
		} else if i < len(result.winningNumbers)+1 {
			result.winningNumbers[i-1] = int(num)
		} else {
			result.numbersOnCard[i-1-len(result.winningNumbers)] = int(num)
		}
	}
	return result
}
