package day7

import "slices"

// ranks represented by ints
// 2-10, Jack=11, Queen=12, King=13, A=14, Joker=-1
type Hand struct {
	cards [5]int
}

func (h Hand) WithJacksAsJokers() Hand {
	for i, card := range h.cards {
		if card == 11 {
			h.cards[i] = -1
		}
	}
	return h // by value
}

// >0 greater, 0 equal, <0 less
func (h Hand) Compare(other Hand) int {
	firstComparison := h.FirstComparisonRule(other)
	if firstComparison != 0 {
		return firstComparison
	} else {
		return h.SecondComparisonRule(other)
	}
}

func (h Hand) FirstComparisonRule(other Hand) int {
	thisPower := h.TypePower()
	otherPower := other.TypePower()
	if thisPower > otherPower {
		return 1
	} else if otherPower > thisPower {
		return -1
	}
	return 0
}

// returns power level of hand based on type
func (h Hand) TypePower() int {
	jokers := 0
	quantities := make([]int, 14)
	for _, card := range h.cards {
		if card == -1 {
			jokers++
		} else {
			quantities[card-1] += 1
		}
	}
	// Sort descending
	slices.SortFunc(quantities, func(a, b int) int {
		if a > b {
			return -1
		} else if a < b {
			return 1
		}
		return 0

	})
	if quantities[0] == 5-jokers {
		// 5 of a kind
		return 6
	}
	if quantities[0] == 4-jokers {
		// 4 of a kind
		return 5
	}
	if quantities[0] >= 3-jokers {
		jokersLeft := jokers - (3 - quantities[0])
		if quantities[1] == 2-jokersLeft {
			// full house
			return 4
		} else {
			// 3 of a kind
			return 3
		}
	}
	if quantities[0] >= 2-jokers {
		jokersLeft := jokers - (2 - quantities[0])
		if quantities[1] == 2-jokersLeft {
			// 2 pair
			return 2
		} else {
			// pair
			return 1
		}
	}
	// nothin'
	return 0
}

func (h Hand) SecondComparisonRule(other Hand) int {
	for i := range h.cards {
		if h.cards[i] > other.cards[i] {
			return 1
		} else if h.cards[i] < other.cards[i] {
			return -1
		}
	}
	return 0
}
