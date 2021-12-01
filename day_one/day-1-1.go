package day_one

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readDayOneInput(filePath string) (result []int) {
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
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, number)
	}

	return result
}

func DayOnePartOne() int {
	input := readDayOneInput("day_one/input_1-1.txt")

	index := 1
	amountOfMeasurements := 0

	for ; index < len(input); index++ {
		if input[index] > input[index-1] {
			amountOfMeasurements++
		}
	}

	return amountOfMeasurements
}
