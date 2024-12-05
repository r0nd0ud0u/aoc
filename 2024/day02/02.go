package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"gitlab.com/Alfred456654/aoc-utils"
)

func main() {
	aoc.CheckEquals("2", step1("test1.txt"))
	fmt.Println(step1("input.txt"))
	aoc.CheckEquals("0", step1("test2.txt"))
}

func step1(filename string) string {
	lines := aoc.ReadInputLines(filename)

	var safe int
	var isSafe bool
	for _, val := range lines {
		splits := strings.Split(val, " ")
		isSafe = false
		var prev int
		var isAscending bool
		if len(splits) > 2 {
			lvl0, err := strconv.Atoi(splits[0])
			if err != nil {
				return "error conv lvl0"
			}
			lvl1, err := strconv.Atoi(splits[1])
			if err != nil {
				return "error conv lvl1"
			}
			diff := lvl1 - lvl0
			if lvl0 == lvl1 || math.Abs(float64(diff)) > 3. {
				isSafe = false
				continue
			}
			if lvl0 < lvl1 {
				isAscending = true
			}
			prev = lvl1
		} else {
			return "split too small"
		}
		for i := 2; i < len(splits); i++ {
			lvl, err := strconv.Atoi(splits[i])
			if err != nil {
				return "error conv"
			}
			diff := lvl - prev
			if diff > 0 && !isAscending || diff < 0 && isAscending || diff == 0 {
				isSafe = false
				break
			}
			if math.Abs(float64(diff)) > 3. {
				isSafe = false
				break
			}
			isSafe = true
			prev = lvl
		}
		if isSafe {
			safe++
		}
	}
	return fmt.Sprintf("%d", safe)
}
