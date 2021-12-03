package day_three

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readInput(filePath string) (result []int, inputLength int) {
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

		inputLength = len(fileLine)

		arg, err := strconv.ParseInt(fileLine, 2, 64)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, int(arg))
	}

	return result, inputLength
}

func DoDayThreePartOne() int {
	inputs, lengthOfInput := readInput("day_three/input_3-1.txt")

	bitCountsGammaRate := make([]int, lengthOfInput)
	bitCountsEpsilonRate := make([]int, lengthOfInput)

	for _, input := range inputs {
		for index := 0; index < lengthOfInput; index++ {
			if input&(1<<index) > 0 {
				bitCountsGammaRate[index] += 1
				bitCountsEpsilonRate[index] -= 1
			} else {
				bitCountsGammaRate[index] -= 1
				bitCountsEpsilonRate[index] += 1
			}
		}
	}

	var gammaRate, epsilonRate int

	for index := 0; index < lengthOfInput; index++ {
		gammaRate += getNormalizedBinary(bitCountsGammaRate[index]) << index
		epsilonRate += getNormalizedBinary(bitCountsEpsilonRate[index]) << index
	}

	return gammaRate * epsilonRate
}

func getNormalizedBinary(number int) int {
	if number > 0 {
		return 1
	}
	return 0
}
