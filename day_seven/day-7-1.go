package day_seven

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput(filePath string) (positions []int) {
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

		splits := strings.Split(fileLine, ",")

		for _, split := range splits {
			pos, err := strconv.Atoi(split)
			if err != nil {
				log.Fatal(err)
			}

			positions = append(positions, pos)
		}
	}

	return positions
}

func getHighestNumber(positions []int) int {
	highest := 0
	for _, position := range positions {
		if position > highest {
			highest = position
		}
	}

	return highest
}

func getFuelConsumptionPartOne(position []int, moveTo int) (lowestConsumption int) {

	for i := 0; i < moveTo; i++ {
		consumption := 0
		for _, position := range position {
			consumption += int(math.Abs(float64(position - i)))
		}

		if consumption < lowestConsumption {
			lowestConsumption = consumption
		} else if lowestConsumption == 0 {
			lowestConsumption = consumption
		}
	}

	return lowestConsumption
}

func DoDaySevenPartOne() int {
	positions := readInput("day_seven/input_7.txt")

	consumption := getFuelConsumptionPartOne(positions, getHighestNumber(positions))

	return consumption
}
