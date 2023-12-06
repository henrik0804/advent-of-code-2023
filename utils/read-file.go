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
