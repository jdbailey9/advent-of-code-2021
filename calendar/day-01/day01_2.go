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

	var depths []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, num)
	}

	var numberOfIncreases = 0

	for i := 1; i < len(depths); i++ {
		if i > 2 && depths[i] > depths[i-3] {
			numberOfIncreases++
		}

	}

	fmt.Println(numberOfIncreases)
}
