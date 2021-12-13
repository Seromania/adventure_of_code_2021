package day_thirteen

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Fold struct {
	x int
	y int
}

type Paper struct {
	folds []Fold
	dots  map[Position]bool
}

func readInput(filePath string) (positions []Position, paper *Paper) {
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

	paper = &Paper{
		folds: []Fold{},
		dots:  make(map[Position]bool),
	}

	scanner := bufio.NewScanner(file)

	hasReachedEmptyLine := false
	for scanner.Scan() {
		fileLine := scanner.Text()

		if fileLine == "" {
			hasReachedEmptyLine = true
			continue
		}

		if hasReachedEmptyLine {
			split := strings.Split(fileLine, "=")
			value, err := strconv.Atoi(split[1])
			if err != nil {
				log.Fatal(err)
			}

			if strings.Contains(split[0], "y") {
				paper.folds = append(paper.folds, Fold{
					x: 0,
					y: value,
				})
			} else {
				paper.folds = append(paper.folds, Fold{
					x: value,
					y: 0,
				})
			}
		} else {
			split := strings.Split(fileLine, ",")
			x, errX := strconv.Atoi(split[0])
			y, errY := strconv.Atoi(split[1])

			if errX != nil {
				log.Fatal(errX)
			}

			if errY != nil {
				log.Fatal(errY)
			}

			positions = append(positions, Position{
				x: x,
				y: y,
			})
		}
	}

	return positions, paper
}

func (p *Paper) fold(positions []Position, partOne bool) {
	toFoldPosition := 0
	if partOne {
		toFoldPosition = 1
	} else {
		toFoldPosition = len(p.folds)
	}

	for _, position := range positions {
		initialX := position.x
		initialY := position.y

		for i := 0; i < toFoldPosition; i++ {
			if p.folds[i].x != 0 && initialX > p.folds[i].x {
				initialX = 2*p.folds[i].x - initialX
			}
			if p.folds[i].y != 0 && initialY > p.folds[i].y {
				initialY = 2*p.folds[i].y - initialY
			}
		}

		p.dots[Position{
			x: initialX,
			y: initialY,
		}] = true
	}
}

func (p *Paper) printDots() {
	for pos := range p.dots {
		fmt.Printf("%d, %d\n", pos.x, pos.y)
	}
}

func (p *Paper) FancyPrint() {
	maxX := -1
	maxY := -1
	for pos := range p.dots {
		if pos.y > maxY {
			maxY = pos.y
		}
		if pos.x > maxX {
			maxX = pos.x
		}
	}

	for y := 0; y < maxY+1; y++ {
		for x := 0; x < maxX+1; x++ {
			if p.dots[Position{
				x: x,
				y: y,
			}] == false {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		println()
	}
}

func DoDayThirteenPartOne() int {
	positions, paper := readInput("day_thirteen/input_13.txt")

	paper.fold(positions, true)
	return len(paper.dots)
}

func DoDayThirteenPartTwo() *Paper {
	positions, paper := readInput("day_thirteen/input_13.txt")

	paper.fold(positions, false)
	return paper
}
