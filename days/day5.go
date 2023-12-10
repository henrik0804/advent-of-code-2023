package days

import (
	"errors"
	"fmt"
	"henrik/advent-of-code-2024/utils"
	"strconv"
	"strings"
)

type RangeEntry struct {
	DestStart   int
	SourceStart int
	RangeLenth  int
}

func DayFive() {

	sourceMapLines := utils.GetInputFileLines("./days/inputs/day5.txt")
	var sourceMaps [][]RangeEntry
	// sourceMap := make(map[string][]RangeEntry)

	seeds := getSeeds(sourceMapLines[0])

	mapKey := ""
	for _, line := range sourceMapLines {

		split := strings.Split(line, " map:")
		if len(split) > 1 {
			mapKey = split[0]
			sourceMaps = append(sourceMaps, make([]RangeEntry, 0))
			continue
		}
		if mapKey != "" {
			convRange, err := getRangeFromLine(line)
			if err != nil {
				continue
			}
			sourceMaps[len(sourceMaps)-1] = append(sourceMaps[len(sourceMaps)-1], convRange)
			//sourceMap[mapKey] = getMapFromRange(destStart, sourceStart, rangeLen, sourceMap[mapKey])
		}
	}

	nearestLocationP1 := int(^uint(0) >> 1)
	nearestLocation := int(^uint(0) >> 1)
	for index, seed := range seeds {
		location := getLocationFromSeed(seed, sourceMaps)
		if location < nearestLocationP1 {
			nearestLocationP1 = location
		}
		if index%2 != 0 {
			continue
		}
		seedRange := seeds[index+1]
		for i := 0; i < seedRange; i++ {
			location := getLocationFromSeed(seed+i, sourceMaps)
			if location < nearestLocation {
				nearestLocation = location
			}
		}
	}

	fmt.Println("The lowest location number for single seeds is: ", nearestLocationP1)
	fmt.Println("The lowest location number for seed ranges is: ", nearestLocation)
}

func getSeeds(seedsLine string) []int {
	seedsString := strings.Split(seedsLine, "seeds: ")[1]
	seedNumStrings := strings.Split(seedsString, " ")
	var seeds []int

	for _, seedStr := range seedNumStrings {
		seed, err := strconv.Atoi(seedStr)
		if err != nil {
			fmt.Println("Error, can't convert seed string to int")
			continue
		}
		seeds = append(seeds, seed)
	}
	return seeds
}

func getRangeFromLine(line string) (RangeEntry, error) {
	rangeParts := strings.Split(line, " ")
	if len(rangeParts) != 3 {
		return RangeEntry{}, errors.New("Not a range-line")
	}
	destStart, err1 := strconv.Atoi(rangeParts[0])
	sourceStart, err2 := strconv.Atoi(rangeParts[1])
	rangeLen, err3 := strconv.Atoi(rangeParts[2])

	if err1 != nil || err2 != nil || err3 != nil {
		return RangeEntry{}, errors.New("Error, can't convert string to number")
	}

	return RangeEntry{DestStart: destStart, SourceStart: sourceStart, RangeLenth: rangeLen}, nil
}

func getMapFromRange(destStart int, sourceStart int, rangeLen int, rangeMap map[int]int) map[int]int {
	for i := 0; i < rangeLen; i++ {
		rangeMap[sourceStart+i] = destStart + i
	}
	return rangeMap
}

func getLocationFromSeed(seed int, sourceMaps [][]RangeEntry) int {
	source := seed
	for _, sourceMap := range sourceMaps {
		newSource := getDestinationValue(source, sourceMap)
		source = newSource
	}
	return source
}
func getDestinationValue(source int, convRange []RangeEntry) int {
	for _, rangeEntry := range convRange {
		sourceIsInRange := source >= rangeEntry.SourceStart && source <= rangeEntry.SourceStart+rangeEntry.RangeLenth-1
		if sourceIsInRange {
			diff := source - rangeEntry.SourceStart
			return rangeEntry.DestStart + diff
		}
	}
	return source
}

/* func getLocationFromSeed(seed int, sourceMaps map[string]map[int]int) int {

	source := seed
	for _, sourceMap := range sourceMaps {
		newSource, ok := sourceMap[source]
		if ok {
			source = newSource
		}
	}
	return source
} */
