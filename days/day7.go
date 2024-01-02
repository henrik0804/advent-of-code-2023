package days

import (
	"errors"
	"fmt"
	"henrik/advent-of-code-2024/utils"
	"sort"
	"strconv"
	"strings"
)

type handType int

const (
	fiveOak handType = iota
	fourOak
	fullHouse
	threeOak
	twoPair
	pair
	highCard
)

type hand struct {
	cards    []rune
	bid      int
	handType handType
}

func DaySeven() {
	inputLines := utils.GetInputFileLines("./days/inputs/day7.txt")
	hands := getHandsFromInput(inputLines)

	sort.Slice(hands, func(i, j int) bool {
		isBetter, err := hands[i].compareHands(&hands[j])
		if err != nil {
			fmt.Println(err)
		}
		return !isBetter
	})
	winnings := 0
	for index, hand := range hands {
		winnings += (index + 1) * hand.bid
	}
	fmt.Println("The total winnings from all hands are: ", winnings)
}

func getHandsFromInput(inp []string) []hand {
	var hands []hand
	for _, line := range inp {
		split := strings.Split(line, " ")

		bid, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err)
		}

		cards := []rune(split[0])
		hType, err := getHandType(cards)
		if err != nil {
			fmt.Println(err)
		}
		hands = append(hands, hand{cards: cards, bid: bid, handType: hType})
	}
	return hands
}

func countRuneOccurrences(runes []rune) []int {
	countMap := make(map[rune]int)

	for _, r := range runes {
		countMap[r]++
	}

	counts := make([]int, len(countMap))

	i := 0
	for _, count := range countMap {
		counts[i] = count
		i++
	}

	return counts
}

func getHandType(runes []rune) (handType, error) {
	counts := countRuneOccurrences(runes)
	if len(counts) == 5 {
		return highCard, nil
	}
	if len(counts) == 4 {
		return pair, nil
	}
	if len(counts) == 1 {
		return fiveOak, nil
	}
	if len(counts) == 2 {
		for _, count := range counts {
			if count == 4 {
				return fourOak, nil
			}
		}
		return fullHouse, nil
	}
	if len(counts) == 3 {
		for _, count := range counts {
			if count == 3 {
				return threeOak, nil
			}
		}
		return twoPair, nil
	}

	return 0, errors.New("No valid Hand")
}

var cardRanks = map[rune]int{
	'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9,
	'9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1,
}

func (h *hand) compareHands(compH *hand) (bool, error) {
	if h.handType != compH.handType {
		return h.handType < compH.handType, nil
	}

	for i := 0; i < len(h.cards) && i < len(compH.cards); i++ {
		rank1 := cardRanks[h.cards[i]]
		rank2 := cardRanks[compH.cards[i]]

		if rank1 != rank2 {
			return rank1 > rank2, nil
		}
	}

	return false, errors.New("Invalid Hand, both are equal")
}
