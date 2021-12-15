package day_10

import (
	"sort"
)

func findAndRemoveCorruptedLines(lines []string) (incompleteLines []string) {
	var lineStack []string

	for _, line := range lines {
		addLine := true

		for i, charInt := range line {
			char := string(charInt)

			if len(lineStack) != 0 && i != 0 {
				if isCharClosing(char) {
					lastCharOfStack := lineStack[len(lineStack)-1]
					lineStack = lineStack[:len(lineStack)-1]

					if !checkIfCorrespondingClosure(lastCharOfStack, char) {
						addLine = false
						break
					}
				} else {
					lineStack = append(lineStack, char)
				}
			} else {
				lineStack = append(lineStack, char)
			}
		}

		if addLine {
			incompleteLines = append(incompleteLines, line)
		}
	}

	return incompleteLines
}

func getCharPoint2(char string) int {
	switch char {
	case "(":
		return 1
	case "[":
		return 2
	case "{":
		return 3
	case "<":
		return 4
	default:
		return 0
	}
}

func findMissingClosingAndCalculateScore(line string) int {
	var lineStack stack

	for i, charInt := range line {
		char := string(charInt)

		if i != 0 {
			if isCharClosing(char) {
				lineStack = lineStack[:len(lineStack)-1]
			} else {
				lineStack = append(lineStack, char)
			}
		} else {
			lineStack = append(lineStack, char)
		}
	}

	score := 0
	for i := len(lineStack) - 1; i >= 0; i-- {
		score = (score * 5) + getCharPoint2(lineStack[i])
	}

	return score
}

func DoDayTenPartTwo() int {
	lines := readInput("day_10/input_ten.txt")

	incompleteLines := findAndRemoveCorruptedLines(lines)

	var scores []int

	for _, line := range incompleteLines {
		score := findMissingClosingAndCalculateScore(line)
		scores = append(scores, score)
	}

	sort.Ints(scores)
	return scores[(len(scores) / 2)]
}
