package main

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/Alfred456654/aoc-utils"
)

type Elem int

const (
	Game Elem = iota
	id
)

const (
	B = "blue"
	G = "green"
	R = "red"
)

var EnumColors = []string{B, G, R}

func step1(filename string, nbG, nbB, nbR int) int {

	lines := aoc.ReadInputLines(filename)
	result := 0
	mapColorsRef := map[string]int{B: nbB, G: nbG, R: nbR}
	for _, line := range lines {
		items := strings.Split(line, " ")
		idNb, errId := strconv.Atoi(items[id][0 : len(items[id])-1])

		isOk := true
		for i := 2; i < len(items)-1; i = i + 2 {
			nb, err := strconv.Atoi(items[i])
			colorStr := items[i+1][0 : len(items[i+1])-1]
			if err == nil && isInEnum(colorStr) && nb > mapColorsRef[colorStr] {
				isOk = false
				break
			}
		}
		if errId == nil && isOk {
			result += idNb
		}
	}

	return result
}

func step2(filename string) int {

	lines := aoc.ReadInputLines(filename)
	result := 0

	for _, line := range lines {
		items := strings.Split(line, " ")
		mapMinColors := map[string]int{B: 0, G: 0, R: 0}
		for i := 2; i < len(items)-1; i = i + 2 {
			nb, err := strconv.Atoi(items[i])
			colorStr := items[i+1][0 : len(items[i+1])-1]
			if err == nil && isInEnum(colorStr) {
				mapMinColors[colorStr] = aoc.Max(mapMinColors[colorStr], nb)
			}
		}
		result += mapMinColors[B] * mapMinColors[G] * mapMinColors[R]
	}

	return result
}

func main() {
	aoc.CheckEquals(8, step1("02_test1.txt", 13, 14, 12))
	fmt.Println(step1("02.txt", 13, 14, 12))
	aoc.CheckEquals(2286, step2("02_test1.txt"))
	fmt.Println(step2("02.txt"))
}

func isInEnum(value string) bool {
	for _, validValue := range EnumColors {
		if value == validValue {
			return true
		}
	}
	return false
}
