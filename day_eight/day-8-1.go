package day_eight

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type digits struct {
	zero  []string
	one   []string
	two   []string
	three []string
	four  []string
	five  []string
	six   []string
	seven []string
	eight []string
	nine  []string
}

type segment struct {
	firstSegment []string
	digits       []string
}

var given_digits = digits{
	zero:  []string{"a", "b", "c", "e", "f", "g"},
	one:   []string{"c", "f"},
	two:   []string{"a", "c", "d", "e", "g"},
	three: []string{"a", "c", "d", "f", "g"},
	four:  []string{"b", "c", "d", "f"},
	five:  []string{"a", "b", "d", "f", "g"},
	six:   []string{"a", "b", "d", "e", "f", "g"},
	seven: []string{"a", "c", "f"},
	eight: []string{"a", "b", "c", "d", "e", "f", "g"},
	nine:  []string{"a", "b", "c", "d", "f", "g"},
}

func readInput(filePath string) (segments []segment) {
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

		splits := strings.Split(fileLine, "|")

		segment := segment{}

		for _, split := range strings.Split(splits[0], " ") {
			if strings.TrimSpace(split) == "" {
				continue
			}
			segment.firstSegment = append(segment.firstSegment, split)
		}

		for _, split := range strings.Split(splits[1], " ") {
			if strings.TrimSpace(split) == "" {
				continue
			}
			segment.digits = append(segment.digits, split)
		}

		segments = append(segments, segment)
	}

	return segments
}

func checkIfArrayHasSegment(digits string, segment []string) bool {
	return len(digits) == len(segment)
}

func doPartOne(digits []string) (hits int) {
	for _, digit := range digits {
		//hasZero := checkIfArrayHasSegment(digit, given_digits.zero)
		hasOne := checkIfArrayHasSegment(digit, given_digits.one)
		//hasTwo := checkIfArrayHasSegment(digit, given_digits.two)
		//hasThree := checkIfArrayHasSegment(digit, given_digits.three)
		hasFour := checkIfArrayHasSegment(digit, given_digits.four)
		//hasFive := checkIfArrayHasSegment(digit, given_digits.five)
		//hasSix := checkIfArrayHasSegment(digit, given_digits.six)
		hasSeven := checkIfArrayHasSegment(digit, given_digits.seven)
		hasEight := checkIfArrayHasSegment(digit, given_digits.eight)
		//hasNine := checkIfArrayHasSegment(digit, given_digits.nine)

		if hasOne || hasFour || hasSeven || hasEight {
			hits++
		}
	}

	return hits
}

func DoDayEightPartOne() int {
	segments := readInput("day_eight/input_8.txt")

	hits := 0
	for _, segment := range segments {
		hits += doPartOne(segment.digits)
	}

	return hits
}
