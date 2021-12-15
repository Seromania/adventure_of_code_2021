package day_5

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Position struct {
	x int
	y int
}

type PositionPair struct {
	pos1 Position
	pos2 Position
}

func transformStringToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		println(err)
		return 0
	}
	return val
}

func readInput(filePath string) (positions []PositionPair) {
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

		//r := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)
		r := regexp.MustCompile(`(\d*),(\d*) -> (\d*),(\d*)`)
		findings := r.FindStringSubmatch(fileLine)

		pos1 := Position{
			x: transformStringToInt(findings[1]),
			y: transformStringToInt(findings[2]),
		}

		pos2 := Position{
			x: transformStringToInt(findings[3]),
			y: transformStringToInt(findings[4]),
		}

		positions = append(positions, PositionPair{
			pos1: pos1,
			pos2: pos2,
		})
	}

	return positions
}

func minInt(first int, second int) int {
	if first > second {
		return second
	}
	return first
}

func maxInt(first int, second int) int {
	if first >= second {
		return first
	}
	return second
}

func iteratePartOne(pos1 Position, pos2 Position, positions map[Position]int) map[Position]int {
	posXSame := true
	start := -1
	end := -1

	if pos1.x == pos2.x {
		posXSame = true
	} else if pos1.y == pos2.y {
		posXSame = false
	} else {
		return positions
	}

	if posXSame {
		start = minInt(pos1.y, pos2.y)
		end = maxInt(pos1.y, pos2.y)
	} else {
		start = minInt(pos1.x, pos2.x)
		end = maxInt(pos1.x, pos2.x)
	}

	for i := start; i <= end; i++ {
		var position Position

		if posXSame {
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

func getAmountOfOverlaps(positions map[Position]int) int {
	overlaps := 0
	for _, value := range positions {
		if value >= 2 {
			overlaps++
		}
	}

	return overlaps
}

func DoDayFivePartOne() int {
	positionPairs := readInput("day_5/input_5.txt")

	lineMap := make(map[Position]int)

	for _, pair := range positionPairs {
		lineMap = iteratePartOne(pair.pos1, pair.pos2, lineMap)
	}

	return getAmountOfOverlaps(lineMap)
}
