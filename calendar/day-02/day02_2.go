package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x   int
	y   int
	aim int
}

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	position := Position{0, 0, 0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		commands := strings.Split(scanner.Text(), " ")

		commandAction := commands[0]
		positionValue, _ := strconv.Atoi(commands[1])

		switch commandAction {
		case string("forward"):
			position.x += positionValue
			position.y += (position.aim * positionValue)
		case string("down"):
			position.aim += positionValue
		case string("up"):
			position.aim -= positionValue
		default:
			panic("bad file")
		}
	}
	fmt.Println(position.x * position.y)
}
