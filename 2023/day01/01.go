package main

import (
	"fmt"
	"sort"
	"strings"

	"gitlab.com/Alfred456654/aoc-utils"
)

var AllNbsMap = map[string]int{"one": 1, "two": 2,
	"three": 3, "four": 4,
	"five": 5, "six": 6,
	"seven": 7, "eight": 8,
	"nine": 9}

func process(filename string, letterCheck bool) int {

	lines := aoc.ReadInputLines(filename)
	result := 0
	for _, i := range lines {
		if len(i) == 0 {
			continue
		}

		mapNbs := make(map[int]int, 0)
		if letterCheck {
			FillMapByLetterNumbers(i, mapNbs)
		}
		for indNb, c := range i {
			if isDigit(byte(c)) {
				mapNbs[indNb] = int(c) - 48
			}
		}

		keys := make([]int, 0, len(mapNbs))
		for k := range mapNbs {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		result += mapNbs[keys[0]]*10 + mapNbs[keys[len(keys)-1]]
	}
	return result
}

func FillMapByLetterNumbers(i string, mapNbs map[int]int) {
	for key := range AllNbsMap {
		ind := strings.Index(i, key)
		ind2 := strings.LastIndex(i, key)
		if ind != -1 {
			mapNbs[ind] = AllNbsMap[key]
		}
		if ind2 != -1 && ind2 != ind {
			mapNbs[ind2] = AllNbsMap[key]
		}
	}
}

func main() {
	aoc.CheckEquals(142, process("01_test1.txt", false))
	aoc.CheckEquals(281, process("01_test2.txt", true))
	fmt.Println(process("01.txt", false))
	fmt.Println(process("01.txt", true))
}

func isDigit(c byte) bool {
	if c >= 49 && c <= 57 {
		return true
	}
	return false
}
