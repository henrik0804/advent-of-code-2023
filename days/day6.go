package days

import (
	"errors"
	"fmt"
	"henrik/advent-of-code-2024/utils"
	"strconv"
	"strings"
)

type race struct {
	time int
	dist int
}

func DaySix() {
	inputLines := utils.GetInputFileLines("./days/inputs/day6.txt")
	races := getRacesFromInput(inputLines)
	singleRace := getSingleRaceFromInput(inputLines)
	fmt.Println(singleRace)

	total := 1
	for _, race := range races {
		wins := race.getWinsCount()
		total *= wins
	}
	fmt.Println("The product of all winning charge Times is :", total)
	fmt.Println("The amount of winning charge Times for the single Race is:", singleRace.getWinsCount())
}

func (r *race) getWinsCount() int {
	min, err := r.calcMinChargeTime()
	if err != nil {
		fmt.Println(err)
	}
	max := r.calcMaxChargeTime(min)
	fmt.Println(min, max)
	winCount := max - min + 1
	return winCount
}

func (r *race) calcMinChargeTime() (int, error) {
	for i := 0; i < r.time; i++ {
		dist := i * (r.time - i)
		if dist > r.dist {
			return i, nil
		}
	}
	return 0, errors.New("Error when calculating charge Time")
}
func (r *race) calcMaxChargeTime(minTime int) int {
	for i := minTime; i < r.time; i++ {
		dist := i * (r.time - i)
		if dist <= r.dist {
			// last winning charge time
			return i - 1
		}
	}
	return 0
}

func getSingleRaceFromInput(inp []string) race {
	timeStr := ""
	distStr := ""
	for _, line := range inp {
		split := strings.Split(line, " ")
		for _, entry := range split {
			num, err := strconv.Atoi(entry)
			if err != nil {
				continue
			}
			if split[0] == "Time:" {
				timeStr += strconv.Itoa(num)
			} else {
				distStr += strconv.Itoa(num)
			}
		}
	}
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		fmt.Println(err)
	}
	dist, err := strconv.Atoi(distStr)
	if err != nil {
		fmt.Println(err)
	}
	return race{time: time, dist: dist}
}

func getRacesFromInput(inp []string) []race {
	var races []race

	for _, line := range inp {
		split := strings.Split(line, " ")
		count := 0
		for _, entry := range split {
			num, err := strconv.Atoi(entry)
			if err != nil {
				continue
			}
			if split[0] == "Time:" {
				races = append(races, race{time: num})
			} else {
				races[count].dist = num
			}
			count++
		}
	}
	return races
}
