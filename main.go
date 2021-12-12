package main

import (
	"adventofcode_2021/day_eight"
	"adventofcode_2021/day_eleven"
	"adventofcode_2021/day_five"
	"adventofcode_2021/day_four"
	"adventofcode_2021/day_nine"
	"adventofcode_2021/day_one"
	"adventofcode_2021/day_seven"
	"adventofcode_2021/day_six"
	"adventofcode_2021/day_ten"
	"adventofcode_2021/day_three"
	"adventofcode_2021/day_twelve"
	"adventofcode_2021/day_two"
	"fmt"
	"time"
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

	runDaySevenPartOne()
	runDaySevenPartTwo()

	runDayEightPartOne()
	runDayEightPartTwo()

	runDayNinePartOne()
	runDayNinePartTwo()

	runDayTenPartOne()
	runDayTenPartTwo()

	runDayElevenPartOne()
	runDayElevenPartTwo()

	runDayTwelvePartOne()
	runDayTwelvePartTwo()
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

func runDaySevenPartOne() {
	result := day_seven.DoDaySevenPartOne()
	fmt.Printf("Result for 7-1: %d\n", result)
}

func runDaySevenPartTwo() {
	result := day_seven.DoDaySevenPartTwo()
	fmt.Printf("Result for 7-2: %d\n", result)
}

func runDayEightPartOne() {
	result := day_eight.DoDayEightPartOne()
	fmt.Printf("Result for 8-1: %d\n", result)
}

func runDayEightPartTwo() {
	result := day_eight.DoDayEightPartTwo()
	fmt.Printf("Result for 8-2: %d\n", result)
}

func runDayNinePartOne() {
	result := day_nine.DoDayNinePartOne()
	fmt.Printf("Result for 9-1: %d\n", result)
}

func runDayNinePartTwo() {
	result := day_nine.DoDayNinePartTwo()
	fmt.Printf("Result for 9-2: %d\n", result)
}

func runDayTenPartOne() {
	result := day_ten.DoDayTenPartOne()
	fmt.Printf("Result for 10-1: %d\n", result)
}

func runDayTenPartTwo() {
	start := time.Now()
	result := day_ten.DoDayTenPartTwo()
	fmt.Printf("Result for 10-2: %d - in %s\n", result, time.Since(start))
}

func runDayElevenPartOne() {
	result := day_eleven.DoDayElevenPartOne()
	fmt.Printf("Result for 11-1: %d\n", result)
}

func runDayElevenPartTwo() {
	result := day_eleven.DoDayElevenPartTwo()
	fmt.Printf("Result for 11-2: %d\n", result)
}

func runDayTwelvePartOne() {
	result := day_twelve.DoDayTwelvePartOne()
	fmt.Printf("Result for 12-1: %d\n", result)
}

func runDayTwelvePartTwo() {
	result := day_twelve.DoDayTwelvePartTwo()
	fmt.Printf("Result for 12-2: %d\n", result)
}
