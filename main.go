package main

import (
	"adventofcode_2021/day_five"
	"adventofcode_2021/day_four"
	"adventofcode_2021/day_one"
	"adventofcode_2021/day_six"
	"adventofcode_2021/day_three"
	"adventofcode_2021/day_two"
	"fmt"
)

func main() {
	runDayOnePartOne()
	runDayOnePartTwo()

	runDayTwoPartOne()
	runDayTwoPartTwo()

	runDayThreePartOne()
	runDayThreePartTwo()

	runDayFourPartOne()
	runDayFourPartTwo()

	runDayFivePartOne()
	runDayFivePartTwo()

	runDaySixPartOne()
	runDaySixPartTwo()
}

func runDayOnePartOne() {
	result := day_one.DayOnePartOne()
	fmt.Printf("Result for 1-1: %d\n", result)
}

func runDayOnePartTwo() {
	result := day_one.DoDayOnePartTwo()
	fmt.Printf("Result for 1-2: %d\n", result)
}

func runDayTwoPartOne() {
	result := day_two.DoDayTwoPartOne()
	fmt.Printf("Result for 2-1: %d\n", result)
}

func runDayTwoPartTwo() {
	result := day_two.DoDayTwoPartTwo()
	fmt.Printf("Result for 2-2: %d\n", result)
}

func runDayThreePartOne() {
	result := day_three.DoDayThreePartOne()
	fmt.Printf("Result for 3-1: %d\n", result)
}

func runDayThreePartTwo() {
	result := day_three.DoDayThreePartTwo()
	fmt.Printf("Result for 3-2: %d\n", result)
}

func runDayFourPartOne() {
	result := day_four.DoDayFourPartOne()
	fmt.Printf("Result for 4-1: %d\n", result)
}

func runDayFourPartTwo() {
	result := day_four.DoDayFourPartTwo()
	fmt.Printf("Result for 4-2: %d\n", result)
}

func runDayFivePartOne() {
	result := day_five.DoDayFivePartOne()
	fmt.Printf("Result for 5-1: %d\n", result)
}

func runDayFivePartTwo() {
	result := day_five.DoDayFivePartTwo()
	fmt.Printf("Result for 5-2: %d\n", result)
}

func runDaySixPartOne() {
	result := day_six.DoDaySixPartOne()
	fmt.Printf("Result for 6-1: %d\n", result)
}

func runDaySixPartTwo() {
	result := day_six.DoDaySixPartTwo()
	fmt.Printf("Result for 6-2: %d\n", result)
}
