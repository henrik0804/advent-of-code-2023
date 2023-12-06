package days

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"henrik/advent-of-code-2024/utils"
	"unicode"
)

type charPosition struct {
	lineNumber int
	charPos    int
}

var inputLines []string = utils.GetInputFileLines("./days/inputs/day3.txt")

/* var inputLines = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
} */

func DayThree() {
	totalPartNumberSum := 0
	totalGearRatio := 0

	for index := range inputLines {
		linePartNumberSum, lineGearRatio := getLinePartNumberSum(index)
		totalPartNumberSum += linePartNumberSum
		totalGearRatio += lineGearRatio
	}

	fmt.Println("The sum of all Part Numbers is ", totalPartNumberSum)
	fmt.Println("The sum of all Gear Ratios is ", totalGearRatio)
}

func getLinePartNumberSum(lineIndex int) (int, int) {
	lineSum := 0
	if lineIndex >= len(inputLines) {
		log.Fatal("Error, Index is out of array bounds")
	}

	charIsNumber := false
	charIndex := 0
	line := []rune(inputLines[lineIndex])

	numberString := ""
	var digitIndices []int

	lineGearRatio := 0

	for charIndex < len(line) {
		char := line[charIndex]
		charIsNumber = isNumber(char)

		if isStarSymbol(char) {
			lineGearRatio += getGearRatioValue(lineIndex, charIndex)
		}

		if charIsNumber {
			numberString += string(char)
			digitIndices = append(digitIndices, charIndex)
		} else {
			charIndex++
			continue
		}

		endOfLine := charIndex == len(line)-1
		nextCharIsNoNumberOrEOL := endOfLine || !isNumber(line[charIndex+1])

		if charIsNumber && (endOfLine || nextCharIsNoNumberOrEOL) {
			lineSum += getPartNumberValue(numberString, lineIndex, digitIndices)
			numberString = ""
			digitIndices = nil

		}

		charIndex++

	}

	return lineSum, lineGearRatio
}

func getPartNumberValue(numberString string, lineIndex int, digitIndices []int) int {
	possiblePartNumber, err := strconv.Atoi(numberString)
	if err != nil {
		log.Fatal("Error, can't convert number string to number")
	}

	hasAdjacentSymbol := false
	for _, digitIndex := range digitIndices {
		if hasAdjacentSymbol {
			return possiblePartNumber
		}

		prevIndex := digitIndex
		if digitIndex > 0 {
			prevIndex = digitIndex - 1
		}

		currentLine := []rune(inputLines[lineIndex])
		nextIndex := digitIndex
		if digitIndex < len(currentLine)-1 {
			nextIndex = digitIndex + 1
		}

		if lineIndex > 0 && !hasAdjacentSymbol {
			prevLine := []rune(inputLines[lineIndex-1])
			hasAdjacentSymbol = isSymbol(prevLine[prevIndex]) || isSymbol(prevLine[digitIndex]) || isSymbol(prevLine[nextIndex])
		}

		if !hasAdjacentSymbol {
			hasAdjacentSymbol = isSymbol(currentLine[prevIndex]) || isSymbol(currentLine[digitIndex]) || isSymbol(currentLine[nextIndex])
		}

		if !hasAdjacentSymbol && lineIndex < len(inputLines[lineIndex])-1 {
			nextLine := []rune(inputLines[lineIndex+1])
			hasAdjacentSymbol = isSymbol(nextLine[prevIndex]) || isSymbol(nextLine[digitIndex]) || isSymbol(nextLine[nextIndex])
		}

	}

	if hasAdjacentSymbol {
		return possiblePartNumber
	}
	return 0
}

func getGearRatioValue(lineIndex int, gearIndex int) int {
	numPositions, err := getAdjNumberIndicesIfGear(lineIndex, gearIndex)
	if err != nil {
		return 0
	}

	num1, num2 := getAdjecentNumbers(numPositions)

	return num1 * num2
}

func getAdjNumberIndicesIfGear(lineIndex int, gearIndex int) ([]charPosition, error) {

	adjacentPartNumberCount := 0

	var numberPositions []charPosition

	currentLine := []rune(inputLines[lineIndex])

	prevIndex := gearIndex
	if gearIndex > 0 {
		prevIndex = gearIndex - 1

		if isNumber(currentLine[prevIndex]) {
			adjacentPartNumberCount++
			numberPositions = append(numberPositions, charPosition{lineIndex, prevIndex})
		}
	}

	nextIndex := gearIndex
	if gearIndex < len(currentLine)-1 {
		nextIndex = gearIndex + 1

		if isNumber(currentLine[nextIndex]) {
			adjacentPartNumberCount++
			numberPositions = append(numberPositions, charPosition{lineIndex, nextIndex})
		}
	}

	if lineIndex > 0 {
		prevLine := []rune(inputLines[lineIndex-1])
		belowIsNumber := isNumber(prevLine[gearIndex])
		if belowIsNumber {
			adjacentPartNumberCount++
			numberPositions = append(numberPositions, charPosition{lineIndex - 1, gearIndex})
		}

		checkPrev := !belowIsNumber && gearIndex > 0
		if checkPrev && isNumber(prevLine[prevIndex]) {
			adjacentPartNumberCount++
			numberPositions = append(numberPositions, charPosition{lineIndex - 1, prevIndex})
		}

		checkNext := !belowIsNumber && gearIndex < len(currentLine)-1
		if checkNext && isNumber(prevLine[nextIndex]) {
			adjacentPartNumberCount++
			numberPositions = append(numberPositions, charPosition{lineIndex - 1, nextIndex})
		}
	}

	if lineIndex < len(inputLines[lineIndex])-1 {
		nextLine := []rune(inputLines[lineIndex+1])
		aboveIsNumber := isNumber(nextLine[gearIndex])
		if aboveIsNumber {
			adjacentPartNumberCount++
			numberPositions = append(numberPositions, charPosition{lineIndex + 1, gearIndex})
		}

		checkPrev := !aboveIsNumber && gearIndex > 0
		if checkPrev && isNumber(nextLine[prevIndex]) {
			adjacentPartNumberCount++
			numberPositions = append(numberPositions, charPosition{lineIndex + 1, prevIndex})
		}

		checkNext := !aboveIsNumber && gearIndex < len(currentLine)-1
		if checkNext && isNumber(nextLine[nextIndex]) {
			adjacentPartNumberCount++
			numberPositions = append(numberPositions, charPosition{lineIndex + 1, nextIndex})
		}
	}

	if adjacentPartNumberCount == 2 {
		return numberPositions, nil
	}
	return nil, errors.New("star symbol is no gear")

}

func getAdjecentNumbers(numberPositions []charPosition) (int, int) {

	var num1 int
	var num2 int
	for _, position := range numberPositions {

		initialDigit := []rune(inputLines[position.lineNumber])[position.charPos]
		numberString := string(initialDigit)

		prevCharIsNumber := true
		nextCharIsNumber := true

		lineNum := position.lineNumber
		charNumBackwards := position.charPos
		charNumForwards := position.charPos

		for prevCharIsNumber {
			prevChar := getPrevChar(charPosition{lineNum, charNumBackwards})
			prevCharIsNumber = isNumber(prevChar)
			charNumBackwards--

			if prevCharIsNumber {
				numberString = string(prevChar) + numberString
			}
		}
		for nextCharIsNumber {
			nextChar := getNextChar(charPosition{lineNum, charNumForwards})
			nextCharIsNumber = isNumber(nextChar)
			charNumForwards++

			if nextCharIsNumber {
				numberString += string(nextChar)
			}
		}

		number, err := strconv.Atoi(numberString)
		if err != nil {
			log.Fatal(err)
		}
		if num1 == 0 {
			num1 = number
		} else {
			num2 = number
		}
	}

	return num1, num2
}

func getNextChar(currentChar charPosition) rune {
	if currentChar.charPos >= len(inputLines[currentChar.lineNumber])-1 {
		return rune('.')
	}
	return []rune(inputLines[currentChar.lineNumber])[currentChar.charPos+1]
}

func getPrevChar(currentChar charPosition) rune {
	if currentChar.charPos <= 0 {
		return rune('.')
	}
	return []rune(inputLines[currentChar.lineNumber])[currentChar.charPos-1]
}

func isNumber(char rune) bool {
	return unicode.IsDigit(char)
}

func isDot(char rune) bool {
	return string(char) == "."
}

func isSymbol(char rune) bool {
	return !isNumber(char) && !isDot(char)
}

func isStarSymbol(char rune) bool {
	return string(char) == "*"
}
