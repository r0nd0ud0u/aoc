package main

import (
	"fmt"
	"strconv"

	"gitlab.com/Alfred456654/aoc-utils"
)

func step1(filename string) int {
	lines := aoc.ReadInputLines(filename)
	result := 0
	for _, i := range lines {
		if len(i) == 0 {
			continue
		}
		strNb := string(i[0]) + string(i[len(i)-1])
		i, _ := strconv.Atoi(strNb)
		result += i
	}
	return result
}

func step2(filename string) int {
}

func main() {
	fmt.Println(step1("01_test.txt"))
	//fmt.Println(step1("01.txt"))
	//fmt.Println(step2("01_test.txt"))
	//fmt.Println(step2("01.txt"))
}
