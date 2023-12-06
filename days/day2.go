package days

import (
	"fmt"
	"henrik/advent-of-code-2024/utils"
	"log"
	"strconv"
	"strings"
)

func DayTwo() {
	inputLines := utils.GetInputFileLines("./days/inputs/day2.txt")

	/* inputLines := []string{
		"game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	} */
	gameIdSum := 0
	totalCubePower := 0
	for index, line := range inputLines {
		gameId := index + 1

		gameRecord := strings.Split(line, ": ")
		if len(gameRecord) == 1 {
			fmt.Println("error, there are no records for this game")
			continue
		}

		drawSets := strings.Split(gameRecord[1], "; ")

		if len(drawSets) == 1 {
			fmt.Println("error, there are no sets in this record")
			continue
		}

		allSetsValid := true
		gameCubePower := getGameCubePower(gameRecord[1])
		for _, drawSet := range drawSets {
			allSetsValid = allSetsValid && getDrawSetIsValid(drawSet)
		}

		if allSetsValid {
			gameIdSum += gameId
		}
		totalCubePower += gameCubePower

	}
	fmt.Println("The sum of all Id's from successful games is ", gameIdSum, " (part 1)")
	fmt.Println("The total cube-power is ", totalCubePower, " (part 2)")

}

func getDrawSetIsValid(setString string) bool {
	set := strings.Split(setString, ", ")
	allDrawsAreValid := true
	for _, draw := range set {
		drawComponents := strings.Split(draw, " ")
		if len(drawComponents) != 2 {
			fmt.Println("error, the draw has the wrong format")
			continue
		}

		drawCount, err := strconv.Atoi(drawComponents[0])
		if err != nil {
			log.Fatal(err)
		}

		allDrawsAreValid = allDrawsAreValid && getDrawIsValid(drawCount, drawComponents[1])
	}
	return allDrawsAreValid
}

func getDrawIsValid(count int, color string) bool {
	const (
		maxCountRed   = 12
		maxCountGreen = 13
		maxCountBlue  = 14
	)

	switch color {
	case "red":
		return count <= maxCountRed
	case "green":
		return count <= maxCountGreen
	case "blue":
		return count <= maxCountBlue
	default:
		return false
	}
}

func getGameCubePower(gameStrings string) int {
	minRedCount := 0
	minGreenCount := 0
	minBlueCount := 0

	drawSets := strings.Split(gameStrings, "; ")

	if len(drawSets) == 1 {
		fmt.Println("error, there are no sets in this record")
	}

	for _, drawSet := range drawSets {
		redCount, greenCount, blueCount := getSetDrawCounts(drawSet)

		if redCount > minRedCount {
			minRedCount = redCount
		}
		if greenCount > minGreenCount {
			minGreenCount = greenCount
		}
		if blueCount > minBlueCount {
			minBlueCount = blueCount
		}
	}
	return minRedCount * minGreenCount * minBlueCount
}

func getSetDrawCounts(setString string) (int, int, int) {
	set := strings.Split(setString, ", ")

	redCount := 0
	greenCount := 0
	blueCount := 0

	for _, draw := range set {
		drawComponents := strings.Split(draw, " ")
		if len(drawComponents) != 2 {
			fmt.Println("error, the draw has the wrong format")
			continue
		}

		drawCount, err := strconv.Atoi(drawComponents[0])
		if err != nil {
			log.Fatal(err)
		}

		drawColor := drawComponents[1]
		switch drawColor {
		case "red":
			redCount = drawCount
		case "green":
			greenCount = drawCount

		case "blue":
			blueCount = drawCount
		}

	}

	return redCount, greenCount, blueCount

}
