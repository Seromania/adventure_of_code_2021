package day_twelve

func checkIfAmountOfVisitIsTwo(path CavePath, caveId string) bool {
	return getAmountOfCaveInPath(path, caveId) >= 2
}

func getAmountOfCaveInPath(path CavePath, caveId string) (amount int) {
	for _, path := range path.path {
		if path == caveId {
			amount += 1
		}
	}

	return amount
}

func (cm *CaveMap) canVisitCaveAgain(caveId string, currentPath CavePath) bool {
	if caveId == "start" || caveId == "end" {
		return false
	}
	if cm.caveMap[caveId].isBig {
		return true
	}

	smallCaveWithMoreThanOneVisitExist := false
	for _, path := range currentPath.path {
		if cm.caveMap[path].isBig {
			continue
		}
		if checkIfAmountOfVisitIsTwo(currentPath, path) {
			smallCaveWithMoreThanOneVisitExist = true
		}
	}

	wantedCaveVisits := getAmountOfCaveInPath(currentPath, caveId)

	if smallCaveWithMoreThanOneVisitExist && wantedCaveVisits == 0 {
		return true
	} else if smallCaveWithMoreThanOneVisitExist && wantedCaveVisits > 0 {
		return false
	}

	return wantedCaveVisits < 2
}

func (cm *CaveMap) findPath2(positionId string, formerHistory CavePath) {
	cavePath := CavePath{
		path:      []string{},
		endsInEnd: false,
	}

	if len(formerHistory.path) > 0 {
		cavePath.path = append(cavePath.path, formerHistory.path...)
	}

	if contains(cavePath.path, "end") {
		cm.paths = append(cm.paths, cavePath)
		return
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
			if connectedCave == "end" {
				if contains(cavePath.path, "end") {
					return
				}
				cavePath.endsInEnd = true
				cavePath.path = append(cavePath.path, connectedCave)
				cm.paths = append(cm.paths, cavePath)
				break
			}
			if !cm.canVisitCaveAgain(connectedCave, cavePath) {
				continue
			}

			cm.findPath2(connectedCave, cavePath)
		}
	}
}

func DoDayTwelvePartTwo() int {
	caveMap := readInput("day_twelve/input_12.txt")

	caveMap.findPath2("start", CavePath{})
	return caveMap.getAmountOfPathsWithEnd()
}
