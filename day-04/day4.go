package day4

import (
	"fmt"
	"os"

	"github.com/golang-collections/collections/set"
)

func Part1(f *os.File) {
	cards := parseInput(f)
	fmt.Printf("Total Points: %d\n", totalPoints(cards))
}

func Part2(f *os.File) {
	cards := parseInput(f)
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
		totalPoints += 1 << (winningNumbers - 1)
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
