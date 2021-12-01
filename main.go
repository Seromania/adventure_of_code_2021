package main

import (
	"adventofcode_2021/day_one"
	"fmt"
)

func main() {
	runDayOnePartOne()
	runDayOnePartTwo()
}

func runDayOnePartOne() {
	result := day_one.DayOnePartOne()
	fmt.Printf("Result for 1-1: %d\n", result)
}

func runDayOnePartTwo() {
	result := day_one.DoDayOnePartTwo()
	fmt.Printf("Result for 1-2: %d\n", result)
}
