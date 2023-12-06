package days

import (
	"fmt"
	"henrik/advent-of-code-2024/utils"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func DayOne() {
	input := utils.GetInputFileLines("./days/inputs/day1.txt")

	sum := 0
	for _, line := range input {
		firstDigit, lastDigit := getNumberStringFromLine(line)
		sum += createNumberFromDigits(firstDigit, lastDigit)
	}
	fmt.Println(sum)
}

func getNumberStringFromLine(line string) (string, string) {

	numStringOptions := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var firstNumString string
	var lowIndex int16 = int16(len(line))

	var secondNumString string
	var highIndex int16

	for _, num := range numStringOptions {
		firstOccurance := strings.Index(line, num)
		lastOccurance := strings.LastIndex(line, num)
		if firstOccurance == -1 {
			continue
		}

		if firstOccurance <= int(lowIndex) {
			lowIndex = int16(firstOccurance)
			firstNumString = num
		}

		if lastOccurance >= int(highIndex) {
			highIndex = int16(lastOccurance)
			secondNumString = num
		}
	}

	return firstNumString, secondNumString
}

// this was for part one, here the number could be only digits 0 ... 9
func getDigitsFromLine(line string) (string, string) {

	var firstDigitAsString string
	var lastDigitAsString string
	for _, char := range line {
		if unicode.IsDigit(char) {
			if firstDigitAsString == "" {
				firstDigitAsString = string(char)
			} else {
				lastDigitAsString = string(char)
			}

		}
	}

	if lastDigitAsString == "" {
		lastDigitAsString = firstDigitAsString
	}

	return firstDigitAsString, lastDigitAsString
}

func createNumberFromDigits(firstDigit string, secondDigit string) int {

	firstDigit = getNumericString(firstDigit)
	secondDigit = getNumericString(secondDigit)

	combinedNumber, err := strconv.Atoi(firstDigit + secondDigit)
	if err != nil {
		log.Fatal(err)
	}

	return combinedNumber
}

func getNumericString(digit string) string {
	switch digit {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return digit
	}
}
