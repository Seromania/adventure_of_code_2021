package day_9

import (
	"sort"
)

type position struct {
	x int
	y int
}

func getLowestPoints2(lavaTubes [][]int) (lowestPointsPosition []position) {
	for posY, lavaTubesX := range lavaTubes {
		for posX := range lavaTubesX {
			if checkAdjacentsLocations(lavaTubes, posY, posX) {
				lowestPointsPosition = append(lowestPointsPosition, position{
					x: posX,
					y: posY,
				})
			}
		}
	}

	return lowestPointsPosition
}

func findHighestThreeAndCalculate(basinSize []int) int {
	sort.Ints(basinSize)
	return basinSize[len(basinSize)-1:][0] * basinSize[len(basinSize)-2 : len(basinSize)-1][0] * basinSize[len(basinSize)-3 : len(basinSize)-1][0]
}

func runThroughTubes(currentPosition position, lavaTubes [][]int, visited [][]bool) int {
	if visited[currentPosition.y][currentPosition.x] {
		return 0
	}

	currentValue := lavaTubes[currentPosition.y][currentPosition.x]

	if currentValue == 9 {
		return 0
	}

	visited[currentPosition.y][currentPosition.x] = true

	size := 1

	if currentPosition.y > 0 && lavaTubes[currentPosition.y-1][currentPosition.x] > currentValue {
		size += runThroughTubes(position{
			x: currentPosition.x,
			y: currentPosition.y - 1,
		}, lavaTubes, visited)
	}
	if currentPosition.y < len(lavaTubes)-1 && lavaTubes[currentPosition.y+1][currentPosition.x] > currentValue {
		size += runThroughTubes(position{
			x: currentPosition.x,
			y: currentPosition.y + 1,
		}, lavaTubes, visited)
	}
	if currentPosition.x > 0 && lavaTubes[currentPosition.y][currentPosition.x-1] > currentValue {
		size += runThroughTubes(position{
			x: currentPosition.x - 1,
			y: currentPosition.y,
		}, lavaTubes, visited)
	}
	if currentPosition.x < len(lavaTubes[0])-1 && lavaTubes[currentPosition.y][currentPosition.x+1] > currentValue {
		size += runThroughTubes(position{
			x: currentPosition.x + 1,
			y: currentPosition.y,
		}, lavaTubes, visited)
	}

	return size
}

func DoDayNinePartTwo() int {
	lavaTubes := readInput("day_9/input_9.txt")

	var visited [][]bool
	visited = make([][]bool, len(lavaTubes))
	for i := range visited {
		visited[i] = make([]bool, len(lavaTubes[i]))
	}

	lowestPointsPosition := getLowestPoints2(lavaTubes)
	var basinSize []int
	for _, lowestPoint := range lowestPointsPosition {
		size := runThroughTubes(position{
			x: lowestPoint.x,
			y: lowestPoint.y,
		}, lavaTubes, visited)

		basinSize = append(basinSize, size)
	}

	return findHighestThreeAndCalculate(basinSize)
}
