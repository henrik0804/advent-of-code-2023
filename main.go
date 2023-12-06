package main

import (
	"bufio"
	"fmt"
	"henrik/advent-of-code-2024/days"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Print("Please enter the Advent calendar door number you want to run: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}
	input = strings.TrimSuffix(input, "\n")

	switch input {
	case "1":
		days.DayOne()
	case "2":
		days.DayTwo()
	}

}