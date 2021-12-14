package day_fourteen

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

func readInput(filePath string) (input string, polymerTemplate map[string]string) {
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

	polymerTemplate = make(map[string]string)

	scanner := bufio.NewScanner(file)

	hasReachedEmptyLine := false
	for scanner.Scan() {
		fileLine := scanner.Text()

		if fileLine == "" {
			hasReachedEmptyLine = true
			continue
		}

		if hasReachedEmptyLine {
			split := strings.Split(fileLine, " ")
			polymerTemplate[split[0]] = split[2]
		} else {
			input = fileLine
		}
	}

	return input, polymerTemplate
}

func calculate(charCountMap map[uint8]uint64) uint64 {
	var leastCommon uint64
	var mostCommon uint64
	leastCommon = math.MaxUint64
	mostCommon = 0

	for _, value := range charCountMap {
		if value > mostCommon {
			mostCommon = value
			continue
		}
		if value < leastCommon {
			leastCommon = value
		}
	}

	return mostCommon - leastCommon
}

func step(pairCount map[string]uint64, polymerTemplate map[string]string) map[string]uint64 {
	newPairCount := make(map[string]uint64)

	for pair, count := range pairCount {
		polyValue, ok := polymerTemplate[pair]
		if !ok {
			panic(pair)
		}

		firstPairChar := string(pair[0])
		secondPairChar := string(pair[1])

		newPairCount[firstPairChar+polyValue] += count
		newPairCount[polyValue+secondPairChar] += count
	}

	return newPairCount
}

func inputToStartingPairCount(input string) map[string]uint64 {
	returnValue := make(map[string]uint64)

	for i := 0; i < len(input)-1; i++ {
		slice := input[i : i+2]
		returnValue[slice] += 1
	}

	return returnValue
}

func GetCharCounts(input string, counts map[string]uint64) map[uint8]uint64 {
	returnCount := make(map[uint8]uint64)
	for pairString, count := range counts {
		if len(pairString) != 2 {
			panic(pairString)
		}
		b := pairString[1]
		returnCount[b] += count
	}

	returnCount[input[0]] += 1
	return returnCount
}

func DoDayFourteenPartOne() uint64 {
	input, polymerTemplate := readInput("day_fourteen/input_14.txt")

	startingPairCount := inputToStartingPairCount(input)
	for i := 0; i < 10; i++ {
		startingPairCount = step(startingPairCount, polymerTemplate)
	}

	counts := GetCharCounts(input, startingPairCount)

	return calculate(counts)
}

func DoDayFourteenPartTwo() uint64 {
	input, polymerTemplate := readInput("day_fourteen/input_14.txt")

	startingPairCount := inputToStartingPairCount(input)
	for i := 0; i < 40; i++ {
		startingPairCount = step(startingPairCount, polymerTemplate)
	}

	counts := GetCharCounts(input, startingPairCount)

	return calculate(counts)
}
