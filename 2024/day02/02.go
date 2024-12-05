package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"gitlab.com/Alfred456654/aoc-utils"
)

func main() {
	aoc.CheckEquals("2", step1("test1.txt"))
	fmt.Println(step1("input.txt"))
	aoc.CheckEquals("4", step2("test2.txt"))
	aoc.CheckEquals("1", step2("test3.txt"))
	fmt.Println(step2("input.txt"))
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

type lvls struct {
	line []int
	cnt  int
}

func step2(filename string) string {
	lines := aoc.ReadInputLines(filename)
	newLines := make([][]int, 0)
	for _, val := range lines {
		splits := strings.Split(val, " ")
		newArr := make([]int, 0)
		for _, val := range splits {
			lvl, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return "error conv lvl"
			}
			newArr = append(newArr, int(lvl))
		}
		newLines = append(newLines, newArr)
	}
	var cnt int
	for _, line := range newLines {
		if isSafe(line) {
			cnt++
			continue
		}

		for i := 0; i < len(line); i++ {
			tmp := slices.Clone(line)
			// check lvl by lvl
			tmp = append(tmp[:i], tmp[i+1:]...)
			if isSafe(tmp) {
				cnt++
				break
			}
		}
	}
	return fmt.Sprintf("%d", cnt)
}

func isSafe(line []int) bool {
	isAscending := line[1] > line[0]
	for i := 1; i < len(line); i++ {
		diff := math.Abs(float64(line[i]) - float64(line[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}
		if (line[i] > line[i-1] && isAscending) || (line[i] < line[i-1] && !isAscending) {
			continue
		}
		return false
	}
	return true
}
