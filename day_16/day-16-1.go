package day_16

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func convertHexStringToBinaryString(hexString string) string {
	binaryString := ""

	for _, char := range hexString {
		switch string(char) {
		case "0":
			binaryString += "0000"
			continue

		case "1":
			binaryString += "0001"
			continue

		case "2":
			binaryString += "0010"
			continue

		case "3":
			binaryString += "0011"
			continue

		case "4":
			binaryString += "0100"
			continue

		case "5":
			binaryString += "0101"
			continue

		case "6":
			binaryString += "0110"
			continue

		case "7":
			binaryString += "0111"
			continue

		case "8":
			binaryString += "1000"
			continue

		case "9":
			binaryString += "1001"
			continue

		case "A":
			binaryString += "1010"
			continue

		case "B":
			binaryString += "1011"
			continue

		case "C":
			binaryString += "1100"
			continue

		case "D":
			binaryString += "1101"
			continue

		case "E":
			binaryString += "1110"
			continue

		case "F":
			binaryString += "1111"
			continue
		}
	}

	return binaryString
}

type BinaryStream struct {
	binaries      string
	stringPointer int
}

func (bs *BinaryStream) get(amount int) (binary string) {
	newPointer := bs.stringPointer + amount

	if newPointer > len(bs.binaries) {
		log.Fatal("New Pointer above length of binary")
	}

	binary = bs.binaries[bs.stringPointer:newPointer]

	bs.stringPointer = newPointer
	return binary
}

func (bs *BinaryStream) seek(amount int) (binary string) {
	return bs.binaries[bs.stringPointer : bs.stringPointer+amount]
}

func (bs *BinaryStream) hasLeft(amount int) bool {
	return bs.stringPointer+amount <= len(bs.binaries)
}

func (bs *BinaryStream) atEnd() bool {
	return bs.stringPointer >= len(bs.binaries)
}

func readInput(filePath string) (binaryStream *BinaryStream) {
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

	binaryStream = &BinaryStream{
		binaries:      "",
		stringPointer: 0,
	}

	scanner := bufio.NewScanner(file)

	hexString := ""
	for scanner.Scan() {
		fileLine := scanner.Text()
		hexString += fileLine
	}

	binaryStream.binaries = convertHexStringToBinaryString(hexString)

	for {
		if len(binaryStream.binaries)%4 == 0 {
			break
		}
		binaryStream.binaries = "0" + binaryStream.binaries
	}

	return binaryStream
}

type Packet struct {
	version            uint8
	typeId             uint8
	literalValueString string
	lengthTypeId       uint8
	calculatedSum      int64
}

func parseBinaryStream(binaryStream *BinaryStream) (result []Packet, calculatedResult int64) {
	currentPacket := Packet{
		version: 0,
		typeId:  0,
	}

	currentPacket = getPacketVersion(binaryStream, currentPacket)
	currentPacket = getPacketTypeId(binaryStream, currentPacket)

	if currentPacket.typeId == 4 {
		currentPacket = getPacketLiteralValueGroups(binaryStream, currentPacket)
		litValue, err := strconv.ParseInt(currentPacket.literalValueString, 2, 64)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, currentPacket)
		return result, litValue
	} else {
		currentPacket = getPacketLengthTypeId(binaryStream, currentPacket)

		amountOfSubPackets := 0
		calculatedSums := []int64{}

		if currentPacket.lengthTypeId == 0 {
			lengthString := binaryStream.get(15)
			length, err := strconv.ParseInt(lengthString, 2, 64)
			if err != nil {
				log.Fatal(err)
			}

			stopAt := binaryStream.stringPointer + int(length)

			for {
				subPackets, calculatedSum := parseBinaryStream(binaryStream)
				amountOfSubPackets += len(subPackets)
				calculatedSums = append(calculatedSums, calculatedSum)
				result = append(result, subPackets...)

				if binaryStream.stringPointer >= stopAt {
					break
				}
			}
		} else {
			lengthString := binaryStream.get(11)
			length, err := strconv.ParseInt(lengthString, 2, 64)
			if err != nil {
				log.Fatal(err)
			}

			for i := int64(0); i < length; i++ {
				subPackets, calculatedSum := parseBinaryStream(binaryStream)
				amountOfSubPackets += len(subPackets)
				calculatedSums = append(calculatedSums, calculatedSum)
				result = append(result, subPackets...)
			}
		}

		switch currentPacket.typeId {
		case 0:
			sum := int64(0)
			for _, subSum := range calculatedSums {
				sum += subSum
			}

			currentPacket.calculatedSum = sum
			result = append(result, currentPacket)
			return result, sum

		case 1:
			sum := int64(1)
			for _, subSum := range calculatedSums {
				sum *= subSum
			}

			currentPacket.calculatedSum = sum
			result = append(result, currentPacket)
			return result, sum

		case 2:
			minimal := int64(math.MaxInt64)
			for _, subSum := range calculatedSums {
				if subSum < minimal {
					minimal = subSum
				}
			}

			currentPacket.calculatedSum = minimal
			result = append(result, currentPacket)
			return result, minimal

		case 3:
			max := int64(0)
			for _, subSum := range calculatedSums {
				if subSum > max {
					max = subSum
				}
			}

			currentPacket.calculatedSum = max
			result = append(result, currentPacket)
			return result, max

		case 5:
			if len(calculatedSums) != 2 {
				panic("5")
			}

			if calculatedSums[0] > calculatedSums[1] {
				currentPacket.calculatedSum = 1
			} else {
				currentPacket.calculatedSum = 0
			}

			result = append(result, currentPacket)
			return result, currentPacket.calculatedSum

		case 6:
			if len(calculatedSums) != 2 {
				panic("6")
			}

			if calculatedSums[0] < calculatedSums[1] {
				currentPacket.calculatedSum = 1
			} else {
				currentPacket.calculatedSum = 0
			}

			result = append(result, currentPacket)
			return result, currentPacket.calculatedSum

		case 7:
			if len(calculatedSums) != 2 {
				panic("6")
			}

			if calculatedSums[0] == calculatedSums[1] {
				currentPacket.calculatedSum = 1
			} else {
				currentPacket.calculatedSum = 0
			}

			result = append(result, currentPacket)
			return result, currentPacket.calculatedSum
		}
	}

	log.Fatal("Criminal scum, you never should have come here")
	return
}

func getPacketVersion(binaryStream *BinaryStream, packet Packet) Packet {
	versionString := binaryStream.get(3)
	version, err := strconv.ParseInt(versionString, 2, 8)
	if err != nil {
		log.Fatal(err)
	}

	packet.version = uint8(version)
	return packet
}

func getPacketTypeId(binaryStream *BinaryStream, packet Packet) Packet {
	typeIdString := binaryStream.get(3)
	typeId, err := strconv.ParseInt(typeIdString, 2, 8)
	if err != nil {
		log.Fatal(err)
	}

	packet.typeId = uint8(typeId)
	return packet
}

func getPacketLiteralValueGroups(binaryStream *BinaryStream, packet Packet) Packet {
	isDone := false

	for isDone == false {
		group := binaryStream.get(5)
		groupValueString := group[1:5]
		packet.literalValueString += groupValueString

		if group[0:1] == "0" {
			isDone = true
		}
	}

	return packet
}

func getPacketLengthTypeId(binaryStream *BinaryStream, packet Packet) Packet {
	lengthTypeIdString := binaryStream.get(1)
	lengthTypeId, err := strconv.ParseInt(lengthTypeIdString, 2, 8)
	if err != nil {
		log.Fatal(err)
	}

	packet.lengthTypeId = uint8(lengthTypeId)
	return packet
}

func doDay(binaryStream *BinaryStream) (versionSummary int64, sum int64) {
	packets, calculatedResult := parseBinaryStream(binaryStream)

	for _, packet := range packets {
		versionSummary += int64(packet.version)
	}

	return versionSummary, calculatedResult
}

func DoDaySixteen() (versionSummary int64, part2 int64) {
	binaryStream := readInput("day_16/input_16.txt")
	return doDay(binaryStream)
}
