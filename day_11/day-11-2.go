package day_11

func (b *board) step2() bool {
	b.currentStep = b.currentStep + 1

	// 1. Increase all octopusi energy by one
	for y := 0; y < b.highestY; y++ {
		for x := 0; x < b.highestX; x++ {
			octopus :=
				b.octopusi[position{
					y: y,
					x: x,
				}]

			octopus.energy = octopus.energy + 1

			b.octopusi[position{
				y: y,
				x: x,
			}] = octopus
		}
	}

	// 2. Flash the octopus whos got the energy
	for y := 0; y < b.highestY; y++ {
		for x := 0; x < b.highestX; x++ {
			currentPos := position{
				y: y,
				x: x,
			}

			if b.octopusi[currentPos].energy > 9 {
				b.flashOctopus(currentPos)
			}
		}
	}

	for y := 0; y < b.highestY; y++ {
		for x := 0; x < b.highestX; x++ {
			currentPos := position{
				y: y,
				x: x,
			}

			if b.octopusi[currentPos].energy > 9 {
				octopus := b.octopusi[currentPos]
				octopus.energy = 0
				b.octopusi[currentPos] = octopus
			}
		}
	}

	isSynchronised := true
	for y := 0; y < b.highestY; y++ {
		for x := 0; x < b.highestX; x++ {
			currentPos := position{
				y: y,
				x: x,
			}

			if b.octopusi[currentPos].energy != 0 {
				isSynchronised = false
				break
			}
		}

		if !isSynchronised {
			break
		}
	}

	return isSynchronised
}

func DoDayElevenPartTwo() int {
	board := readInput("day_11/input_11.txt")

	for {
		isSynced := board.step2()
		if isSynced {
			break
		}
	}

	return board.currentStep
}
