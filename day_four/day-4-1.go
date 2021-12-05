package day_four

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filePath string) (bingo Bingo) {
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

	isReadingValueDrawn := true

	bingo = Bingo{
		drawnNumbers: []int{},
		bingoBoards:  []*BingoBoard{},
	}

	currentBingoBoard := initBingoBoard()
	currentRow := 0

	for scanner.Scan() {
		fileLine := scanner.Text()

		if isReadingValueDrawn {
			if len(fileLine) == 0 {
				isReadingValueDrawn = false
				continue
			}

			numbersString := strings.Split(fileLine, ",")

			for _, numberStr := range numbersString {
				number, err := strconv.Atoi(numberStr)

				if err != nil {
					log.Fatalf("Couldnt cast number %s", numberStr)
				}

				bingo.drawnNumbers = append(bingo.drawnNumbers, number)
			}
		} else {
			if len(fileLine) == 0 {
				bingo.bingoBoards = append(bingo.bingoBoards, currentBingoBoard)
				currentBingoBoard = initBingoBoard()
				currentRow = 0
				continue
			}

			numberStrings := strings.Split(fileLine, " ")

			var numberValues []int
			for _, numberStr := range numberStrings {
				if numberStr == "" {
					continue
				}

				number, err := strconv.Atoi(numberStr)
				if err != nil {
					log.Fatal(err)
				}
				numberValues = append(numberValues, number)
			}

			currentBingoBoard.setRow(currentRow, numberValues)
			currentRow++
		}
	}

	bingo.bingoBoards = append(bingo.bingoBoards, currentBingoBoard)

	return bingo
}

func DoDayFourPartOne() int {
	bingo := readInput("day_four/input_day-4.txt")

	for index, drawnNumber := range bingo.drawnNumbers {
		for _, board := range bingo.bingoBoards {
			board.checkNumberAndMarkIfFound(drawnNumber)

			if index >= 5 {
				if board.checkIfWon() {
					return board.getSumOfAllUnmarkedFields() * drawnNumber
				}
			}
		}
	}

	return 0
}
