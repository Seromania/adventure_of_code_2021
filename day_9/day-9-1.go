package day_9

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readInput(filePath string) (lavaTubes [][]int) {
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

	line := 0
	for scanner.Scan() {
		fileLine := scanner.Text()

		lavaTubes = append(lavaTubes, []int{})
		lavaTubes[line] = make([]int, len(fileLine))

		for i, numberStr := range fileLine {
			number, err := strconv.Atoi(string(numberStr))
			if err != nil {
				log.Fatal(err)
			}

			lavaTubes[line][i] = number
		}

		line++
	}

	return lavaTubes
}

func checkAdjacentsLocations(lavaTubes [][]int, positionY int, positionX int) bool {
	numberToCheck := lavaTubes[positionY][positionX]

	// Up
	if positionY-1 >= 0 {
		if numberToCheck >= lavaTubes[positionY-1][positionX] {
			return false
		}
	}

	// Left
	if positionX-1 >= 0 {
		if numberToCheck >= lavaTubes[positionY][positionX-1] {
			return false
		}
	}

	// Right
	if positionX+1 < len(lavaTubes[positionY]) {
		if numberToCheck >= lavaTubes[positionY][positionX+1] {
			return false
		}
	}

	// Down
	if positionY+1 < len(lavaTubes) {
		if numberToCheck >= lavaTubes[positionY+1][positionX] {
			return false
		}
	}

	return true
}

func getLowestPoints(lavaTubes [][]int) (lowestPoints []int) {
	for posY, lavaTubesX := range lavaTubes {
		for posX := range lavaTubesX {
			if checkAdjacentsLocations(lavaTubes, posY, posX) {
				lowestPoints = append(lowestPoints, lavaTubes[posY][posX])
			}
		}
	}

	return lowestPoints
}

func DoDayNinePartOne() int {
	lavaTubes := readInput("day_9/input_test.txt")

	lowestPoints := getLowestPoints(lavaTubes)

	riskSum := 0
	for _, lowestPoint := range lowestPoints {
		riskSum += lowestPoint + 1
	}

	return riskSum
}
