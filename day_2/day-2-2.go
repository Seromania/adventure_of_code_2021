package day_2

func (c *Command) HandlePartTwo(horizontalPosition int, depth int, aim int) (int, int, int) {
	switch c.commandName {
	case "forward":
		horizontalPosition += c.argument
		depth += aim * c.argument
		break

	case "down":
		aim += c.argument
		break

	case "up":
		aim -= c.argument
		break
	}

	return horizontalPosition, depth, aim
}

func DoDayTwoPartTwo() int {
	horizontalPosition := 0
	depth := 0
	aim := 0

	commands := readInput("day_2/input_2-1.txt")
	//commands := readInput("day_2/input_test.txt.txt")

	for _, command := range commands {
		horizontalPosition, depth, aim = command.HandlePartTwo(horizontalPosition, depth, aim)
	}

	return horizontalPosition * depth
}
