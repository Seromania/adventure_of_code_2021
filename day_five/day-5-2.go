package day_five

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

	if pos1.x == pos2.x {
		iterateMode = iterateModeSameX
	} else if pos1.y == pos2.y {
		iterateMode = iterateModeSameY
	} else if absInt(pos1.x) == absInt(pos2.x) &&
		absInt(pos1.y) == absInt(pos2.y) {
		iterateMode = iterateModeDiagonal
	} else {
		return positions
	}

	switch iterateMode {
	case iterateModeSameX:
		start = minInt(pos1.y, pos2.y)
		end = maxInt(pos1.y, pos2.y)
		break

	case iterateModeSameY:
		start = minInt(pos1.x, pos2.x)
		end = maxInt(pos1.x, pos2.x)
		break
	}

	for i := start; i <= end; i++ {
		var position Position

		if iterateMode {
			position = Position{
				x: pos1.x,
				y: i,
			}
		} else {
			position = Position{
				x: i,
				y: pos1.y,
			}
		}

		if val, ok := positions[position]; ok {
			positions[position] = val + 1
		} else {
			positions[position] = 1
		}
	}

	return positions
}
