package day_3

func calcBitCount(inputs []int, index int) (bitCount int) {
	for _, input := range inputs {
		if input&(1<<index) > 0 {
			bitCount++
		} else {
			bitCount--
		}
	}

	return bitCount
}

func getMostCommonBit(inputs []int, index int) int {
	bitCount := calcBitCount(inputs, index)

	if bitCount < 0 {
		return 0
	} else {
		return 1
	}
}

func getLeastCommonBit(inputs []int, index int) int {
	bitCount := calcBitCount(inputs, index)

	if bitCount < 0 {
		return 1
	} else {
		return 0
	}
}

func removeNumberByCriteria(inputs []int, index int, shouldBeZero bool) (results []int) {
	for _, input := range inputs {
		if shouldBeZero && input&(1<<index) == 0 ||
			!shouldBeZero && input&(1<<index) > 0 {
			results = append(results, input)
		}
	}

	return results
}

func getOxygenGeneratorRating(inputs []int, index int) (result int) {
	if len(inputs) == 1 {
		return inputs[0]
	}

	var shouldBeZero bool

	if getMostCommonBit(inputs, index) == 0 {
		shouldBeZero = true
	} else {
		shouldBeZero = false
	}

	inputs = removeNumberByCriteria(inputs, index, shouldBeZero)

	return getOxygenGeneratorRating(inputs, index-1)
}

func getCO2ScrubberRating(inputs []int, index int) (result int) {
	if len(inputs) == 1 {
		return inputs[0]
	}

	var shouldBeZero bool

	if getLeastCommonBit(inputs, index) == 0 {
		shouldBeZero = true
	} else {
		shouldBeZero = false
	}

	inputs = removeNumberByCriteria(inputs, index, shouldBeZero)

	return getCO2ScrubberRating(inputs, index-1)
}

func DoDayThreePartTwo() int {
	inputs, lengthOfInput := readInput("day_3/input_3-1.txt")

	oxygenGeneratorRating := getOxygenGeneratorRating(inputs, lengthOfInput-1)
	co2ScrubberRating := getCO2ScrubberRating(inputs, lengthOfInput-1)

	return oxygenGeneratorRating * co2ScrubberRating
}
