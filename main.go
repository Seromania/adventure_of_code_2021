package main

import (
	"adventofcode_2021/day_1"
	"adventofcode_2021/day_10"
	"adventofcode_2021/day_11"
	"adventofcode_2021/day_12"
	"adventofcode_2021/day_13"
	"adventofcode_2021/day_14"
	"adventofcode_2021/day_15"
	"adventofcode_2021/day_2"
	"adventofcode_2021/day_3"
	"adventofcode_2021/day_4"
	"adventofcode_2021/day_5"
	"adventofcode_2021/day_6"
	"adventofcode_2021/day_7"
	"adventofcode_2021/day_8"
	"adventofcode_2021/day_9"
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

	runDayThirteenPartOne()
	runDayThirteenPartTwo()

	runDayFourteenPartOne()
	runDayFourteenPartTwo()

	runDayFifteenPartOne()
	runDayFifteenPartTwo()
}

func runDayOnePartOne() {
	result := day_1.DayOnePartOne()
	fmt.Printf("Result for 1-1: %d\n", result)
}

func runDayOnePartTwo() {
	result := day_1.DoDayOnePartTwo()
	fmt.Printf("Result for 1-2: %d\n", result)
}

func runDayTwoPartOne() {
	result := day_2.DoDayTwoPartOne()
	fmt.Printf("Result for 2-1: %d\n", result)
}

func runDayTwoPartTwo() {
	result := day_2.DoDayTwoPartTwo()
	fmt.Printf("Result for 2-2: %d\n", result)
}

func runDayThreePartOne() {
	result := day_3.DoDayThreePartOne()
	fmt.Printf("Result for 3-1: %d\n", result)
}

func runDayThreePartTwo() {
	result := day_3.DoDayThreePartTwo()
	fmt.Printf("Result for 3-2: %d\n", result)
}

func runDayFourPartOne() {
	result := day_4.DoDayFourPartOne()
	fmt.Printf("Result for 4-1: %d\n", result)
}

func runDayFourPartTwo() {
	result := day_4.DoDayFourPartTwo()
	fmt.Printf("Result for 4-2: %d\n", result)
}

func runDayFivePartOne() {
	result := day_5.DoDayFivePartOne()
	fmt.Printf("Result for 5-1: %d\n", result)
}

func runDayFivePartTwo() {
	result := day_5.DoDayFivePartTwo()
	fmt.Printf("Result for 5-2: %d\n", result)
}

func runDaySixPartOne() {
	result := day_6.DoDaySixPartOne()
	fmt.Printf("Result for 6-1: %d\n", result)
}

func runDaySixPartTwo() {
	result := day_6.DoDaySixPartTwo()
	fmt.Printf("Result for 6-2: %d\n", result)
}

func runDaySevenPartOne() {
	result := day_7.DoDaySevenPartOne()
	fmt.Printf("Result for 7-1: %d\n", result)
}

func runDaySevenPartTwo() {
	result := day_7.DoDaySevenPartTwo()
	fmt.Printf("Result for 7-2: %d\n", result)
}

func runDayEightPartOne() {
	result := day_8.DoDayEightPartOne()
	fmt.Printf("Result for 8-1: %d\n", result)
}

func runDayEightPartTwo() {
	result := day_8.DoDayEightPartTwo()
	fmt.Printf("Result for 8-2: %d\n", result)
}

func runDayNinePartOne() {
	result := day_9.DoDayNinePartOne()
	fmt.Printf("Result for 9-1: %d\n", result)
}

func runDayNinePartTwo() {
	result := day_9.DoDayNinePartTwo()
	fmt.Printf("Result for 9-2: %d\n", result)
}

func runDayTenPartOne() {
	result := day_10.DoDayTenPartOne()
	fmt.Printf("Result for 10-1: %d\n", result)
}

func runDayTenPartTwo() {
	start := time.Now()
	result := day_10.DoDayTenPartTwo()
	fmt.Printf("Result for 10-2: %d - in %s\n", result, time.Since(start))
}

func runDayElevenPartOne() {
	result := day_11.DoDayElevenPartOne()
	fmt.Printf("Result for 11-1: %d\n", result)
}

func runDayElevenPartTwo() {
	result := day_11.DoDayElevenPartTwo()
	fmt.Printf("Result for 11-2: %d\n", result)
}

func runDayTwelvePartOne() {
	result := day_12.DoDayTwelvePartOne()
	fmt.Printf("Result for 12-1: %d\n", result)
}

func runDayTwelvePartTwo() {
	result := day_12.DoDayTwelvePartTwo()
	fmt.Printf("Result for 12-2: %d\n", result)
}

func runDayThirteenPartOne() {
	result := day_13.DoDayThirteenPartOne()
	fmt.Printf("Result for 13-1: %d\n", result)
}

func runDayThirteenPartTwo() {
	result := day_13.DoDayThirteenPartTwo()
	fmt.Printf("Result for 13-2:\n")
	result.FancyPrint()
}

func runDayFourteenPartOne() {
	result := day_14.DoDayFourteenPartOne()
	fmt.Printf("Result for 14-1:%d\n", result)
}

func runDayFourteenPartTwo() {
	result := day_14.DoDayFourteenPartTwo()
	fmt.Printf("Result for 14-2:%d\n", result)
}

func runDayFifteenPartOne() {
	result := day_15.DoDayFifteenPartOne()
	fmt.Printf("Result for 15-1: %d\n", result)
}

func runDayFifteenPartTwo() {
	result := day_15.DoDayFifteenPartTwo()
	fmt.Printf("Result for 15-2: %d\n", result)
}
