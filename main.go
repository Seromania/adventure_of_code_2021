package main

import (
	"adventofcode_2021/day_one"
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
