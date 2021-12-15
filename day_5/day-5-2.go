package day_5

type iterateMode int

const (
	iterateModeSameX iterateMode = iota
	iterateModeSameY
	iterateModeDiagonal
)

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func iteratePartTwo(pos1 Position, pos2 Position, positions map[Position]int) map[Position]int {
	iterateMode := iterateModeSameX
	start := -1
	end := -1
	firstIterator := 1

	secondStart := -1
	secondEnd := -1
	secondIterator := 1

	if pos1.x == pos2.x {
		iterateMode = iterateModeSameX
	} else if pos1.y == pos2.y {
		iterateMode = iterateModeSameY
	} else if absInt(pos1.x-pos2.x) == absInt(pos1.y-pos2.y) {
		iterateMode = iterateModeDiagonal
	} else {
		return positions
	}

	switch iterateMode {
	case iterateModeSameX:
		start = minInt(pos1.y, pos2.y)
		end = maxInt(pos1.y, pos2.y)
		firstIterator = 1
		break

	case iterateModeSameY:
		start = minInt(pos1.x, pos2.x)
		end = maxInt(pos1.x, pos2.x)
		firstIterator = 1
		break

	case iterateModeDiagonal:
		start = pos1.x
		end = pos2.x
		secondStart = pos1.y
		secondEnd = pos2.y

		if start > end {
			firstIterator = -1
		} else {
			firstIterator = 1
		}

		if secondStart > secondEnd {
			secondIterator = -1
		} else {
			secondIterator = 1
		}

		break
	}

	stepEnd := absInt(end - start)

	for i := 0; i <= stepEnd; i++ {
		var position Position

		switch iterateMode {
		case iterateModeSameX:
			position = Position{
				x: pos1.x,
				y: start,
			}
			break

		case iterateModeSameY:
			position = Position{
				x: start,
				y: pos1.y,
			}
			break

		case iterateModeDiagonal:
			position = Position{
				x: start,
				y: secondStart,
			}
			break
		}

		if val, ok := positions[position]; ok {
			positions[position] = val + 1
		} else {
			positions[position] = 1
		}

		start += firstIterator
		secondStart += secondIterator
	}

	return positions
}

func DoDayFivePartTwo() int {
	positionPairs := readInput("day_5/input_5.txt")

	lineMap := make(map[Position]int)

	for _, pair := range positionPairs {
		lineMap = iteratePartTwo(pair.pos1, pair.pos2, lineMap)
	}

	return getAmountOfOverlaps(lineMap)
}
