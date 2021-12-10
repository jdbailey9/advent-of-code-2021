package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	var binaryNumbers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		binaryNumbers = append(binaryNumbers, scanner.Text())
	}

	gammaRate, epsilonRate := 0, 0
	ones := map[int]int{}

	for _, number := range binaryNumbers {
		for pos := len(number) - 1; pos >= 0; pos-- {
			index, _ := strconv.Atoi(string(number[pos]))
			if index == 1 {
				ones[len(number)-1-pos]++
			}
		}
	}
	for pos, count := range ones {
		if count > (len(binaryNumbers) - count) {
			gammaRate += 1 << pos
		} else {
			epsilonRate += 1 << pos
		}
	}

	fmt.Println(gammaRate * epsilonRate)
}
