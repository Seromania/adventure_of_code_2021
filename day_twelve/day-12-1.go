package day_twelve

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Cave struct {
	id          string
	isBig       bool
	connectedTo []string
}

type CavePath struct {
	path      []string
	endsInEnd bool
}

type CaveMap struct {
	caveMap map[string]Cave
	paths   []CavePath
}

func contains(stringArray []string, string string) bool {
	for _, v := range stringArray {
		if v == string {
			return true
		}
	}

	return false
}

func (cm *CaveMap) findPath(positionId string, formerHistory CavePath) {
	cavePath := CavePath{
		path:      []string{},
		endsInEnd: false,
	}

	if len(formerHistory.path) > 0 {
		cavePath.path = formerHistory.path
	}

	cave := cm.caveMap[positionId]

	cavePathPathLen := len(cavePath.path)
	if cavePathPathLen == 0 || cavePathPathLen > 0 && cavePath.path[cavePathPathLen-1] != cave.id {
		cavePath.path = append(cavePath.path, cave.id)
	}

	if cave.id == "end" {
		cavePath.endsInEnd = true
	} else {
		for _, connectedCave := range cave.connectedTo {
			// if connectedCave is not big and was already in path, continue
			if !cm.caveMap[connectedCave].isBig && contains(cavePath.path, connectedCave) {
				continue
			}

			cm.findPath(connectedCave, cavePath)
		}
	}

	cm.paths = append(cm.paths, cavePath)
}

func (cm *CaveMap) printPaths() {
	for _, path := range cm.paths {
		for _, caveId := range path.path {
			fmt.Printf("%s,", caveId)
		}
		println()
	}
}

func (cm *CaveMap) printPathsWithEnd() {
	for _, path := range cm.paths {
		if !path.endsInEnd {
			continue
		}
		for _, caveId := range path.path {
			fmt.Printf("%s,", caveId)
		}
		println()
	}
}

func (cm *CaveMap) getAmountOfPathsWithEnd() (amount int) {
	for _, path := range cm.paths {
		if !path.endsInEnd {
			continue
		}
		amount = amount + 1
	}

	return amount
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func readInput(filePath string) (caveMap *CaveMap) {
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

	caveMap = &CaveMap{
		caveMap: make(map[string]Cave),
		paths:   []CavePath{},
	}

	for scanner.Scan() {
		fileLine := scanner.Text()

		splits := strings.Split(fileLine, "-")

		caveOne := caveMap.caveMap[splits[0]]
		caveTwo := caveMap.caveMap[splits[1]]

		caveOne.id = splits[0]
		if caveOne.id != splits[1] && !contains(caveOne.connectedTo, splits[1]) {
			caveOne.connectedTo = append(caveOne.connectedTo, splits[1])
		}

		if IsUpper(caveOne.id) {
			caveOne.isBig = true
		} else {
			caveOne.isBig = false
		}

		caveTwo.id = splits[1]
		if caveTwo.id != splits[0] && !contains(caveTwo.connectedTo, splits[0]) {
			caveTwo.connectedTo = append(caveTwo.connectedTo, splits[0])
		}

		if IsUpper(caveTwo.id) {
			caveTwo.isBig = true
		} else {
			caveTwo.isBig = false
		}

		caveMap.caveMap[caveOne.id] = caveOne
		caveMap.caveMap[caveTwo.id] = caveTwo
	}

	return caveMap
}

func DoDayTwelvePartOne() int {
	caveMap := readInput("day_twelve/input_12.txt")

	caveMap.findPath("start", CavePath{})
	//caveMap.printPaths()
	caveMap.printPathsWithEnd()
	return caveMap.getAmountOfPathsWithEnd()
}
