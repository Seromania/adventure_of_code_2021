package day_2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	commandName string
	argument    int
}

func (c *Command) Handle(horizontalPosition int, depth int) (int, int) {
	switch c.commandName {
	case "forward":
		horizontalPosition += c.argument
		break

	case "down":
		depth += c.argument
		break

	case "up":
		depth -= c.argument
		break
	}

	return horizontalPosition, depth
}

func readInput(filePath string) (result []Command) {
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
		lineSplit := strings.Split(fileLine, " ")

		commandStr := lineSplit[0]
		argumentStr := lineSplit[1]
		arg, err := strconv.Atoi(argumentStr)
		if err != nil {
			log.Fatal(err)
		}

		command := Command{
			commandName: commandStr,
			argument:    arg,
		}

		result = append(result, command)
	}

	return result
}

func DoDayTwoPartOne() int {
	horizontalPosition := 0
	depth := 0

	commands := readInput("day_2/input_2-1.txt")

	for _, command := range commands {
		horizontalPosition, depth = command.Handle(horizontalPosition, depth)
	}

	return horizontalPosition * depth
}
