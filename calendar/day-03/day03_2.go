package main

import (
	"bufio"
	"fmt"
	"math"
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

	getRating := func(compFunc func(int, int) bool) int {
		lines := make([]string, len(binaryNumbers))
		copy(lines, binaryNumbers)
		for b := 0; b < len(binaryNumbers[0]); b++ {
			newLines := []string{}
			ones, target := 0, 0
			for _, l := range lines {
				num, _ := strconv.Atoi(string(l[b]))
				ones += num
			}
			if compFunc(ones, len(lines)-ones) {
				target++
			}
			for _, l := range lines {
				num, _ := strconv.Atoi(string(l[b]))
				if num == target {
					newLines = append(newLines, l)
				}
			}
			if len(newLines) == 1 {
				num, _ := strconv.Atoi(newLines[0])
				return binaryToDecimal(num)
			}
			lines = newLines
		}
		panic("hmmmmmmmmmmmmmmmmmmmmm")
	}

	oxygenRating := getRating(func(i1, i2 int) bool { return i1 >= i2 })
	co2Rating := getRating(func(i1, i2 int) bool { return i1 < i2 })
	fmt.Println(oxygenRating * co2Rating)
}

func binaryToDecimal(num int) int {
	var remainder int
	index := 0
	decimalNum := 0
	for num != 0 {
		remainder = num % 10
		num = num / 10
		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum
}
