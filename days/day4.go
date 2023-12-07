package days

import (
	"errors"
	"fmt"
	"henrik/advent-of-code-2024/utils"
	"strconv"
	"strings"
)

type ticket struct {
	winningNumbers []int
	ticketNumbers  []int
}

func DayFour() {
	var scratchCards []string = utils.GetInputFileLines("./days/inputs/day4.txt")
	/* scratchCards := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	} */

	tickets := getTickets(scratchCards)

	totalPoints := 0

	for _, ticket := range tickets {
		winCount := getWinningNumberMatchCount(ticket)

		if winCount > 1 {
			totalPoints += getPointNumber(winCount)
		} else {
			totalPoints += winCount
		}
	}

	fmt.Println("The total Point-Number from all Scratchcards is: ", totalPoints)

}

func getTickets(scratchCards []string) []ticket {
	var tickets []ticket
	for _, line := range scratchCards {
		cardContent := strings.Split(line, ":")
		if len(cardContent) != 2 {
			fmt.Println("Error, the Input string has the wrong format")
		}
		contentSplit := strings.Split(cardContent[1], "|")

		if len(contentSplit) != 2 {
			fmt.Println("Error, the Input string has the wrong format")
		}

		var winningNumbers []int
		var ticketNumbers []int

		for _, winNum := range strings.Split(contentSplit[0], " ") {
			num, err := convertToNumber(winNum)
			if err == nil {
				winningNumbers = append(winningNumbers, num)
			}
		}

		for _, ticketNum := range strings.Split(contentSplit[1], " ") {
			num, err := convertToNumber(ticketNum)
			if err == nil {
				ticketNumbers = append(ticketNumbers, num)
			}
		}
		tickets = append(tickets, ticket{winningNumbers: winningNumbers, ticketNumbers: ticketNumbers})
	}
	return tickets
}

func getWinningNumberMatchCount(ticket ticket) int {
	count := 0

	for _, num := range ticket.ticketNumbers {
		numIncluded := utils.ArrayContains(ticket.winningNumbers, num)
		if numIncluded {
			count++
		}
	}

	return count
}

func getPointNumber(matchCount int) int {
	points := 1
	for i := 0; i < matchCount-1; i++ {
		points *= 2
	}
	return points
}

func convertToNumber(numString string) (int, error) {
	digitsOnly := strings.TrimSpace(numString)

	if len(digitsOnly) == 0 {
		return 0, errors.New("String is empty")
	}

	num, err := strconv.Atoi(digitsOnly)
	if err != nil {
		fmt.Println(err)
	}

	return num, nil
}
