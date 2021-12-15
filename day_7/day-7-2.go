package day_7

import "math"

func getConsumption(difference int) (consumption int) {
	for i := 0; i < difference; i++ {
		consumption += i + 1
	}

	return consumption
}

func getFuelConsumptionPartTwo(position []int, moveTo int) (lowestConsumption int) {
	for i := 0; i < moveTo; i++ {
		consumption := 0
		for _, position := range position {
			consumption += getConsumption(int(math.Abs(float64(position - i))))
		}

		if consumption < lowestConsumption {
			lowestConsumption = consumption
		} else if lowestConsumption == 0 {
			lowestConsumption = consumption
		}
	}

	return lowestConsumption
}

func DoDaySevenPartTwo() int {
	positions := readInput("day_7/input_7.txt")
	result := getFuelConsumptionPartTwo(positions, getHighestNumber(positions))

	return result
}
