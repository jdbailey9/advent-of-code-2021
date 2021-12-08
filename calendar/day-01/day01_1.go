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

	var numberOfIncreases int

	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			numberOfIncreases++
		}
	}

	fmt.Println(numberOfIncreases)
}
