package main

import (
	"aoc2022/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

func main() {
	//fmt.Println(step1("05_test.txt"))
	fmt.Println(step("05.txt", true))
	fmt.Println(step("05.txt", false))
}

func step(filename string, step1 bool) string {
	result := ""
	lines := utils.ReadInputLinesWithoutTrim(filename)
	isStack := true
	nbStacks := 0
	var stackTable [][]string

	for _, line := range lines {
		if strings.Contains(line, "move") {
			isStack = false
		}
		if line == "" {
			continue
		}

		if isStack {
			if nbStacks == 0 {
				nbStacks = (len(line) + 1) / 4
				stackTable = make([][]string, nbStacks)
			}

			for i := 0; i < nbStacks; i++ {
				crate := string(line[4*i+1])
				_, err := strconv.Atoi(crate)
				if err == nil {
					break
				}
				if crate != " " {
					if stackTable[i] == nil {
						stackTable[i] = make([]string, 0)
					}

					stackTable[i] = append(stackTable[i], crate)
				}
			}
		} else {
			matches := re.FindAllStringSubmatch(line, -1)
			number, _ := strconv.Atoi(matches[0][1])
			prevPos, _ := strconv.Atoi(matches[0][2])
			newPos, _ := strconv.Atoi(matches[0][3])
			stackTable = DoMoves(step1, stackTable, number, prevPos-1, newPos-1)
		}
	}

	for _, stack := range stackTable {
		result += stack[0]
	}
	return result
}

func DoMoves(step1 bool, stacks [][]string, nb, prev, next int) [][]string {
	if step1 {
		for i := 0; i < nb; i++ {
			moved := stacks[prev][0]
			stacks[next] = append([]string{moved}, stacks[next]...)
			stacks[prev] = stacks[prev][1:]
		}
	} else {
		moved := make([]string, nb)
		copy(moved, stacks[prev][:nb])

		stacks[prev] = stacks[prev][nb:]
		stacks[next] = append(moved, stacks[next]...)
	}
	return stacks
}
