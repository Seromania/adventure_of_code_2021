package day_15

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/beefsack/go-astar"
)

type Position struct {
	x int
	y int
}

type Cave struct {
	caveRooms map[Position]int
	maxX      int
	maxY      int
}

func readInput(filePath string) (cave *Cave) {
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

	cave = &Cave{
		caveRooms: make(map[Position]int),
		maxX:      0,
		maxY:      0,
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileLine := scanner.Text()
		cave.maxX = len(fileLine)
		for i, char := range fileLine {
			number, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}

			cave.caveRooms[Position{
				x: i,
				y: cave.maxY,
			}] = number
		}

		cave.maxY += 1
	}

	return cave
}

type CaveTile struct {
	cave *Cave

	x int
	y int
}

func IntAbs(n int) int {
	if n < 0 {
		return -1 * n
	}

	return n
}

func ManhattanDistance(x1, y1, x2, y2 int) int {
	return IntAbs(x2-x1) + IntAbs(y2-y1)
}

func (ct CaveTile) PathNeighbors() (results []astar.Pather) {
	for _, delta := range []struct {
		x int
		y int
	}{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		_, ok := ct.cave.caveRooms[Position{
			x: ct.x + delta.x,
			y: ct.y + delta.y,
		}]

		if !ok {
			continue
		}

		results = append(results, CaveTile{cave: ct.cave, x: ct.x + delta.x, y: ct.y + delta.y})
	}

	return results
}

func (ct CaveTile) PathNeighborCost(to astar.Pather) float64 {
	toTile := to.(CaveTile)
	r, _ := ct.cave.caveRooms[Position{
		x: toTile.x,
		y: toTile.y,
	}]

	return float64(r)
}

func (ct CaveTile) PathEstimatedCost(to astar.Pather) float64 {
	other := to.(CaveTile)
	return float64(ManhattanDistance(ct.x, ct.y, other.x, other.y))
}

var _ astar.Pather = CaveTile{}

func DoDayFifteenPartOne() int {
	cave := readInput("day_15/input_15.txt")

	_, distance, found := astar.Path(
		CaveTile{cave: cave, x: 0, y: 0},
		CaveTile{cave: cave, x: cave.maxX - 1, y: cave.maxY - 1})
	if !found {
		panic("no solution")
	}

	return int(distance)
}
