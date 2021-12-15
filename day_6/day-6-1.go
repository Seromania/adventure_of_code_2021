package day_6

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filePath string) (fishes []int) {
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

	for scanner.Scan() {
		fileLine := scanner.Text()

		fishesStr := strings.Split(fileLine, ",")

		for _, fishStr := range fishesStr {
			num, err := strconv.Atoi(fishStr)
			if err != nil {
				log.Fatal(err)
			}

			fishes = append(fishes, num)
		}
	}

	return fishes
}

type Fishes struct {
	zeroDay  int
	oneDay   int
	twoDay   int
	threeDay int
	fourDay  int
	fiveDay  int
	sixDay   int
	sevenDay int
	eightDay int
}

func initializeFishes(fishes []int) *Fishes {
	myFishes := &Fishes{
		zeroDay:  0,
		oneDay:   0,
		twoDay:   0,
		threeDay: 0,
		fourDay:  0,
		fiveDay:  0,
		sixDay:   0,
		sevenDay: 0,
		eightDay: 0,
	}

	for _, fish := range fishes {
		if fish == 0 {
			myFishes.zeroDay++
		} else if fish == 1 {
			myFishes.oneDay++
		} else if fish == 2 {
			myFishes.twoDay++
		} else if fish == 3 {
			myFishes.threeDay++
		} else if fish == 4 {
			myFishes.fourDay++
		} else if fish == 5 {
			myFishes.fiveDay++
		} else if fish == 6 {
			myFishes.sixDay++
		} else if fish == 7 {
			myFishes.sevenDay++
		} else if fish == 8 {
			myFishes.eightDay++
		}
	}

	return myFishes
}

func (fishes *Fishes) tick() {
	zeroDay := fishes.zeroDay
	oneDay := fishes.oneDay
	twoDay := fishes.twoDay
	threeDay := fishes.threeDay
	fourDay := fishes.fourDay
	fiveDay := fishes.fiveDay
	sixDay := fishes.sixDay
	sevenDay := fishes.sevenDay
	eightDay := fishes.eightDay

	fishes.zeroDay = oneDay
	fishes.oneDay = twoDay
	fishes.twoDay = threeDay
	fishes.threeDay = fourDay
	fishes.fourDay = fiveDay
	fishes.fiveDay = sixDay
	fishes.sixDay = sevenDay
	fishes.sevenDay = eightDay
	fishes.eightDay = zeroDay

	fishes.sixDay += zeroDay
}

func tickDays(days int, fishes []int) int {
	fr := initializeFishes(fishes)

	for i := 0; i < days; i++ {
		fr.tick()
	}

	return fr.zeroDay +
		fr.oneDay +
		fr.twoDay +
		fr.threeDay +
		fr.fourDay +
		fr.fiveDay +
		fr.sixDay +
		fr.sevenDay +
		fr.eightDay
}

func DoDaySixPartOne() int {
	fishes := readInput("day_6/input_6.txt")

	return tickDays(80, fishes)
}

func DoDaySixPartTwo() int {
	fishes := readInput("day_6/input_6.txt")

	return tickDays(256, fishes)
}
