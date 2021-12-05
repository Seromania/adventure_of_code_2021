package day_four

import (
	"fmt"
	"log"
)

type BingoField struct {
	number int
	marked bool
}

type BingoBoard struct {
	grid [][]*BingoField
}

func initBingoBoard() (bingoBoard *BingoBoard) {
	bingoFields := make([][]*BingoField, 5)
	for i := range bingoFields {
		bingoFields[i] = make([]*BingoField, 5)
	}

	for i, row := range bingoFields {
		for j := range row {
			bingoFields[i][j] = &BingoField{
				number: 0,
				marked: false,
			}
		}
	}

	bingoBoard = &BingoBoard{
		grid: bingoFields,
	}

	return bingoBoard
}

func (bb *BingoBoard) print() {
	for _, row := range bb.grid {
		for _, column := range row {
			if column.marked {
				fmt.Printf("[%d]\t", column.number)
			} else {
				fmt.Printf("%d\t", column.number)
			}
		}

		fmt.Printf("\n")
	}
}

func (bb *BingoBoard) setRow(row int, values []int) {
	if len(values) != 5 {
		log.Fatal("Setting row with not 5 values")
	}

	bb.grid[row][0].number = values[0]
	bb.grid[row][1].number = values[1]
	bb.grid[row][2].number = values[2]
	bb.grid[row][3].number = values[3]
	bb.grid[row][4].number = values[4]
}

func (bb *BingoBoard) checkNumberAndMarkIfFound(number int) {
	for _, row := range bb.grid {
		for _, column := range row {
			if column.number == number {
				column.marked = true
				return
			}
		}
	}
}

func (bb *BingoBoard) checkIfWon() bool {
	return bb.checkIfRowIsMarked(0) ||
		bb.checkIfRowIsMarked(1) ||
		bb.checkIfRowIsMarked(2) ||
		bb.checkIfRowIsMarked(3) ||
		bb.checkIfRowIsMarked(4) ||
		bb.checkIfColumnIsMarked(0) ||
		bb.checkIfColumnIsMarked(1) ||
		bb.checkIfColumnIsMarked(2) ||
		bb.checkIfColumnIsMarked(3) ||
		bb.checkIfColumnIsMarked(4)
}

func (bb *BingoBoard) checkIfRowIsMarked(row int) bool {
	return bb.grid[row][0].marked == true &&
		bb.grid[row][1].marked == true &&
		bb.grid[row][2].marked == true &&
		bb.grid[row][3].marked == true &&
		bb.grid[row][4].marked == true
}

func (bb *BingoBoard) checkIfColumnIsMarked(column int) bool {
	return bb.grid[0][column].marked == true &&
		bb.grid[1][column].marked == true &&
		bb.grid[2][column].marked == true &&
		bb.grid[3][column].marked == true &&
		bb.grid[4][column].marked == true
}

func (bb *BingoBoard) getSumOfAllUnmarkedFields() int {
	sum := 0
	for _, row := range bb.grid {
		for _, column := range row {
			if !column.marked {
				sum += column.number
			}
		}
	}

	return sum
}

type Bingo struct {
	drawnNumbers []int
	bingoBoards  []*BingoBoard
}
