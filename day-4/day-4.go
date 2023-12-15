package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/golang-collections/collections/set"
	"github.com/mitchellh/iochan"
)

func main() {
	f, _ := os.Open("./input.txt")
	cards := make([]Card, 1)
	for line := range iochan.DelimReader(f, '\n') {
		card := parseCard(line)
		cards = append(cards, card)
	}
	fmt.Printf("Total Points: %d\n", totalPoints(cards))
	fmt.Printf("Total Cards: %d\n", totalCards(cards))
}

// Part 1
func totalPoints(cards []Card) int {
	totalPoints := 0
	for _, card := range cards {
		winningNumbers := card.winningNumbersOnCard()
		if winningNumbers == 0 {
			continue
		}
		totalPoints += int(math.Pow(2, float64(winningNumbers-1)))
	}
	return totalPoints
}

// Part 2
func totalCards(cards []Card) int {
	quantities := make([]int, len(cards))
	for i := range quantities {
		quantities[i] = 1
	}

	for i, card := range cards {
		winningNumbers := card.winningNumbersOnCard()
		for won := 1; won <= winningNumbers; won++ {
			if i+won >= len(quantities) {
				break
			}
			quantities[i+won] += quantities[i]
		}
	}

	totalCards := 0
	for _, quantity := range quantities {
		totalCards += quantity
	}

	return totalCards
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

func (card *Card) winningNumbersOnCard() int {
	winning := set.New()
	for _, winningNum := range card.winningNumbers {
		winning.Insert(winningNum)
	}
	onCard := set.New()
	for _, cardNum := range card.numbersOnCard {
		onCard.Insert(cardNum)
	}

	common := winning.Intersection(onCard)
	return common.Len()
}

type Card struct {
	cardNum        int
	winningNumbers [10]int
	numbersOnCard  [25]int
}
