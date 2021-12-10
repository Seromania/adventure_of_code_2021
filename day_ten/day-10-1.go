package day_ten

import (
	"bufio"
	"log"
	"os"
)

type stack []string

func readInput(filePath string) (lines []string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileLine := scanner.Text()
		lines = append(lines, fileLine)
	}

	return lines
}

func checkIfCorrespondingClosure(formerChar string, currentChar string) bool {
	switch formerChar {
	case "[":
		return currentChar == "]"

	case "(":
		return currentChar == ")"

	case "{":
		return currentChar == "}"

	case "<":
		return currentChar == ">"
	}

	return false
}

func getCharPoint(char string) int {
	switch char {
	case "]":
		return 57

	case ")":
		return 3

	case "}":
		return 1197

	case ">":
		return 25137
	}

	return 0
}

func isCharClosing(char string) bool {
	switch char {
	case "]":
		return true

	case ")":
		return true

	case "}":
		return true

	case ">":
		return true

	default:
		return false
	}
}

func findLineCorruptionReturnPoint(line string) int {
	var lineStack []string

	for i, charInt := range line {
		char := string(charInt)

		if len(lineStack) != 0 && i != 0 {
			if isCharClosing(char) {
				lastCharOfStack := lineStack[len(lineStack)-1]
				lineStack = lineStack[:len(lineStack)-1]

				if !checkIfCorrespondingClosure(lastCharOfStack, char) {
					return getCharPoint(char)
				}
			} else {
				lineStack = append(lineStack, char)
			}
		} else {
			lineStack = append(lineStack, char)
		}
	}

	return 0
}

func DoDayTenPartOne() int {
	lines := readInput("day_ten/input_ten.txt")

	points := 0
	for _, line := range lines {
		points += findLineCorruptionReturnPoint(line)
	}

	return points
}
