package utils

import (
	"bufio"
	"log"
	"os"
)

func GetInputFileLines(filepath string) []string {

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	return fileLines
}

func ArrayContains(arr []int, value int) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}
