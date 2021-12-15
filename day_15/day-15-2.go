package day_15

import (
	"github.com/beefsack/go-astar"
)

func (c *Cave) extendCave() {
	oldMap := c.caveRooms
	newMap := make(map[Position]int)

	oldMaxX := c.maxX
	oldMaxY := c.maxY

	c.maxX = c.maxX * 5
	c.maxY = c.maxY * 5

	for productX := 0; productX < 5; productX++ {
		for productY := 0; productY < 5; productY++ {
			for x := 0; x < oldMaxX; x++ {
				for y := 0; y < oldMaxY; y++ {
					value, _ := oldMap[Position{
						x: x,
						y: y,
					}]

					var formerRisk int
					var ok bool
					if productX > 0 {
						formerRisk, ok = newMap[Position{
							x: oldMaxX*(productX-1) + x,
							y: oldMaxX*productY + y,
						}]
					} else if productY > 0 {
						formerRisk, ok = newMap[Position{
							x: oldMaxX*productX + x,
							y: oldMaxX*(productY-1) + y,
						}]
					}

					if ok {
						value = formerRisk + 1

						if value == 10 {
							value = 1
						}
					}

					newMap[Position{
						x: oldMaxX*productX + x,
						y: oldMaxY*productY + y,
					}] = value
				}
			}
		}
	}

	c.caveRooms = newMap
}

func DoDayFifteenPartTwo() int {
	cave := readInput("day_15/input_15.txt")
	cave.extendCave()

	_, distance, found := astar.Path(
		CaveTile{cave: cave, x: 0, y: 0},
		CaveTile{cave: cave, x: cave.maxX - 1, y: cave.maxY - 1})
	if !found {
		panic("no solution")
	}

	return int(distance)
}
