package day_eleven

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type position struct {
	y int
	x int
}

type octopus struct {
	energy          int
	lastFlashAtStep int
}

type board struct {
	octopusi    map[position]octopus
	currentStep int

	highestY int
	highestX int

	flashes int
}

func (b *board) print() {
	fmt.Printf("==Step: %d==\n", b.currentStep)
	for y := 0; y < b.highestY; y++ {
		for x := 0; x < b.highestX; x++ {
			fmt.Printf("%d", b.octopusi[position{
				y: y,
				x: x,
			}].energy)
		}
		fmt.Printf("\n")
	}
	println()
}

func (b *board) flashOctopus(at position) {
	if b.octopusi[at].energy < 10 || b.octopusi[at].lastFlashAtStep == b.currentStep {
		return
	}

	octopus := b.octopusi[at]
	octopus.lastFlashAtStep = b.currentStep

	b.flashes = b.flashes + 1
	b.octopusi[at] = octopus

	for _, deltaPosition := range [][]int{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}} {
		deltaAt := position{
			y: at.y + deltaPosition[0],
			x: at.x + deltaPosition[1],
		}
		deltaOctopus := b.octopusi[deltaAt]
		if deltaOctopus.lastFlashAtStep == 0 && deltaOctopus.energy == 0 {
			continue
		}

		deltaOctopus.energy = deltaOctopus.energy + 1
		b.octopusi[deltaAt] = deltaOctopus

		if deltaOctopus.lastFlashAtStep == b.currentStep {
			continue
		}

		if deltaOctopus.energy > 9 {
			b.flashOctopus(deltaAt)
		}
	}
}

func (b *board) step() {
	b.currentStep = b.currentStep + 1

	// 1. Increase all octopusi energy by one
	for y := 0; y < b.highestY; y++ {
		for x := 0; x < b.highestX; x++ {
			octopus :=
				b.octopusi[position{
					y: y,
					x: x,
				}]

			octopus.energy = octopus.energy + 1

			b.octopusi[position{
				y: y,
				x: x,
			}] = octopus
		}
	}

	// 2. Flash the octopus whos got the energy
	for y := 0; y < b.highestY; y++ {
		for x := 0; x < b.highestX; x++ {
			currentPos := position{
				y: y,
				x: x,
			}

			if b.octopusi[currentPos].energy > 9 {
				b.flashOctopus(currentPos)
			}
		}
	}

	for y := 0; y < b.highestY; y++ {
		for x := 0; x < b.highestX; x++ {
			currentPos := position{
				y: y,
				x: x,
			}

			if b.octopusi[currentPos].energy > 9 {
				octopus := b.octopusi[currentPos]
				octopus.energy = 0
				b.octopusi[currentPos] = octopus
			}
		}
	}
}

func readInput(filePath string) (returnBoard *board) {
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

	returnBoard = &board{
		octopusi:    make(map[position]octopus),
		currentStep: 0,
		highestY:    0,
		highestX:    0,
	}

	returnBoard.currentStep = 0
	y := 0
	for scanner.Scan() {
		fileLine := scanner.Text()

		for x, charInt := range fileLine {
			level, err := strconv.Atoi(string(charInt))
			if err != nil {
				log.Fatal(err)
			}

			returnBoard.octopusi[position{
				y: y,
				x: x,
			}] = octopus{
				energy:          level,
				lastFlashAtStep: 0,
			}

			returnBoard.highestX = x + 1
		}

		y++
		returnBoard.highestY = y
	}

	return returnBoard
}

func DoDayElevenPartOne() int {
	board := readInput("day_eleven/input_11.txt")

	for i := 0; i < 100; i++ {
		board.step()
	}
	return board.flashes
}
