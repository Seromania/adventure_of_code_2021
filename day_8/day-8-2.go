package day_8

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func checkIfSortedDigitContainsAllSegmentCodeChars(sortedDigit string, segmentCode string) bool {
	for i := 0; i < len(segmentCode); i++ {
		if !strings.Contains(sortedDigit, string(segmentCode[i])) {
			return false
		}
	}
	return true
}

func checkIfSortedDigitIntersectsSegmentCodeCharThreeTimes(sortedDigit string, segmentCode string) bool {
	intersects := 0
	for i := 0; i < len(segmentCode); i++ {
		if strings.Contains(sortedDigit, string(segmentCode[i])) {
			intersects++
		}
	}
	return intersects == 3
}

func getTwoThreeOrFive(segmentCode [4]string, digit string) string {
	runeArray := []rune(digit)
	sort.Sort(sortRuneString(runeArray))
	sortedDigit := string(runeArray)

	if checkIfSortedDigitContainsAllSegmentCodeChars(sortedDigit, segmentCode[2]) {
		return "3"
	}

	if checkIfSortedDigitIntersectsSegmentCodeCharThreeTimes(sortedDigit, segmentCode[1]) {
		return "5"
	}

	return "2"
}

func getZeroSixOrNine(segmentCode [4]string, digit string) string {
	runeArray := []rune(digit)
	sort.Sort(sortRuneString(runeArray))
	sortedDigit := string(runeArray)

	if checkIfSortedDigitContainsAllSegmentCodeChars(sortedDigit, segmentCode[1]) {
		return "9"
	}

	if checkIfSortedDigitContainsAllSegmentCodeChars(sortedDigit, segmentCode[2]) {
		return "0"
	}

	return "6"
}

func getEntriesForOneFourSevenAndEight(segments []string) [4]string {
	one := ""
	four := ""
	seven := ""
	eight := ""

	for _, segment := range segments {
		if checkIfArrayHasSegment(segment, given_digits.one) {
			runeArray := []rune(segment)
			sort.Sort(sortRuneString(runeArray))
			one = string(runeArray)
		} else if checkIfArrayHasSegment(segment, given_digits.four) {
			runeArray := []rune(segment)
			sort.Sort(sortRuneString(runeArray))
			four = string(runeArray)
		} else if checkIfArrayHasSegment(segment, given_digits.seven) {
			runeArray := []rune(segment)
			sort.Sort(sortRuneString(runeArray))
			seven = string(runeArray)
		} else if checkIfArrayHasSegment(segment, given_digits.eight) {
			runeArray := []rune(segment)
			sort.Sort(sortRuneString(runeArray))
			eight = string(runeArray)
		}
	}

	return [4]string{
		one,
		four,
		seven,
		eight,
	}
}

func getSegmentByDigit(segments []string, digits []string) (number int) {
	numberStr := ""

	segmentCode := getEntriesForOneFourSevenAndEight(segments)

	for _, digit := range digits {
		if checkIfArrayHasSegment(digit, given_digits.one) {
			numberStr += "1"
			continue
		}
		if checkIfArrayHasSegment(digit, given_digits.four) {
			numberStr += "4"
			continue
		}
		if checkIfArrayHasSegment(digit, given_digits.seven) {
			numberStr += "7"
			continue
		}
		if checkIfArrayHasSegment(digit, given_digits.eight) {
			numberStr += "8"
			continue
		}

		if len(digit) == 5 {
			numberStr += getTwoThreeOrFive(segmentCode, digit)
		}
		if len(digit) == 6 {
			numberStr += getZeroSixOrNine(segmentCode, digit)
		}
	}

	if numberStr == "" {
		return 0
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		println(numberStr)
		log.Fatal(err)
	}

	return number
}

func DoDayEightPartTwo() int {
	segments := readInput("day_8/input_8.txt")
	hits := 0

	for _, segment := range segments {
		hits += getSegmentByDigit(segment.firstSegment, segment.digits)
	}

	return hits
}
