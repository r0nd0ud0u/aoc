package main

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

// string mapping
const (
	ROCK_1       = "A"
	PAPER_1      = "B"
	SCISSORS_1   = "C"
	ROCK_2       = "X"
	PAPER_2      = "Y"
	SCISSORS_2   = "Z"
	NEED_TO_LOSE = "X"
	NEED_TO_DRAW = "Y"
	NEED_TO_WIN  = "Z"
)

const (
	LOST       = 0
	DRAW       = 3
	WIN        = 6
	V_ROCK     = 1
	V_PAPER    = 2
	V_SCISSORS = 3
)

func main() {
	// expected 15
	fmt.Println(step1("02_test.txt"))
	fmt.Println(step1("02.txt"))
	// expected 12
	fmt.Println(step2("02_test.txt"))
	fmt.Println(step2("02.txt"))
}

func readInput(filename string) [][]string {
	stratLines := utils.ReadInputLines(filename)
	result := make([][]string, len(stratLines))
	for idx, line := range stratLines {
		chunks := strings.Split(line, " ")
		result[idx] = []string{chunks[0], chunks[1]}
	}
	return result
}

func step1(filename string) int {
	lines := readInput(filename)
	result := 0
	for _, line := range lines {
		add := janken(line[0], line[1])
		result = result + add
	}
	return result
}

func step2(filename string) int {
	lines := readInput(filename)
	result := 0
	for _, line := range lines {
		type2 := changeType2ForJanken(line[0], line[1])
		add := janken(line[0], type2)
		result = result + add

	}
	return result
}

func janken(A string, B string) int {
	if A == ROCK_1 && B == ROCK_2 {
		return DRAW + V_ROCK
	}
	if A == ROCK_1 && B == PAPER_2 {
		return WIN + V_PAPER
	}
	if A == ROCK_1 && B == SCISSORS_2 {
		return LOST + V_SCISSORS
	}
	if A == PAPER_1 && B == PAPER_2 {
		return DRAW + V_PAPER
	}
	if A == PAPER_1 && B == SCISSORS_2 {
		return WIN + V_SCISSORS
	}
	if A == PAPER_1 && B == ROCK_2 {
		return LOST + V_ROCK
	}
	if A == SCISSORS_1 && B == ROCK_2 {
		return WIN + V_ROCK
	}
	if A == SCISSORS_1 && B == PAPER_2 {
		return LOST + V_PAPER
	}
	if A == SCISSORS_1 && B == SCISSORS_2 {
		return DRAW + V_SCISSORS
	}

	return LOST
}

func changeType2ForJanken(A string, B string) string {
	if A == ROCK_1 && B == NEED_TO_DRAW {
		return ROCK_2
	}
	if A == ROCK_1 && B == NEED_TO_WIN {
		return PAPER_2
	}
	if A == ROCK_1 && B == NEED_TO_LOSE {
		return SCISSORS_2
	}
	if A == PAPER_1 && B == NEED_TO_DRAW {
		return PAPER_2
	}
	if A == PAPER_1 && B == NEED_TO_WIN {
		return SCISSORS_2
	}
	if A == PAPER_1 && B == NEED_TO_LOSE {
		return ROCK_2
	}
	if A == SCISSORS_1 && B == NEED_TO_WIN {
		return ROCK_2
	}
	if A == SCISSORS_1 && B == NEED_TO_LOSE {
		return PAPER_2
	}
	if A == SCISSORS_1 && B == NEED_TO_DRAW {
		return SCISSORS_2
	}

	return "heheh"
}
