package day_4

func RemoveIndex(s []*BingoBoard, index int) []*BingoBoard {
	ret := make([]*BingoBoard, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func DoDayFourPartTwo() int {
	bingo := readInput("day_4/input_day-4.txt")

	for index, drawnNumber := range bingo.drawnNumbers {
		boardsToRemove := make([]*BingoBoard, 0)

		for _, board := range bingo.bingoBoards {
			board.checkNumberAndMarkIfFound(drawnNumber)

			if index >= 5 {
				if board.checkIfWon() {
					boardsToRemove = append(boardsToRemove, board)

					if len(bingo.bingoBoards) == 1 {
						return board.getSumOfAllUnmarkedFields() * drawnNumber
					}
				}
			}
		}

		if len(boardsToRemove) == 0 {
			continue
		}

		for _, boardToRemove := range boardsToRemove {
			for indexInBingo, boardInBingo := range bingo.bingoBoards {
				if boardToRemove == boardInBingo {
					boardInBingo = nil
					bingo.bingoBoards = RemoveIndex(bingo.bingoBoards, indexInBingo)
					break
				}
			}
		}

	}

	return 0
}
