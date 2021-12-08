package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	position := Position{0, 0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		commands := strings.Split(scanner.Text(), " ")

		commandAction := commands[0]
		positionValue, _ := strconv.Atoi(commands[1])

		switch commandAction {
		case string("forward"):
			position.x += positionValue
		case string("down"):
			position.y += positionValue
		case string("up"):
			position.y -= positionValue
		default:
			panic("bad file")
		}
	}
	fmt.Println(position.x * position.y)
}
