package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadIntsByBlocks(filename string) [][]int {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Errorf("could not read file '%s'", err)
	}
	blocks := strings.Split(string(bytes), "\r\n\r\n")
	result := make([][]int, len(blocks))
	for i, b := range blocks {
		lines := strings.Split(b, "\n")
		result[i] = make([]int, len(lines))
		for j, l := range lines {
			val64, err2 := strconv.ParseInt(strings.TrimSpace(l), 10, 64)
			if err2 != nil {
				fmt.Errorf("cannot parse int '%s'", err2, l)
			}
			result[i][j] = int(val64)
		}
	}
	return result
}

func minandmax(values []int) (int, int) {
	min := values[0] //assign the first element equal to min
	max := values[0] //assign the first element equal to max
	for _, number := range values {
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}
	return min, max
}

func main() {
	result := ReadIntsByBlocks("puzzles/01.txt")
	//result := ReadIntsByBlocks("puzzles/01_test.txt")
	sums := make([]int, len(result))
	for i, nain := range result {
		for _, calories := range nain {
			sums[i] += calories
		}
	}
	_, maxSlice := minandmax(sums)
	fmt.Println(maxSlice)

}
