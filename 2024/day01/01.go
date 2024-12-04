package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"gitlab.com/Alfred456654/aoc-utils"
)

func main() {
	aoc.CheckEquals("11", step1("01_test1.txt", true))
	fmt.Println(step1("01.txt", true))
	aoc.CheckEquals("31", step1("01_test2.txt", false))
	fmt.Println(step1("01.txt", false))
}

func step1(filename string, isStep1 bool) string {
	lines := aoc.ReadInputLines(filename)

	col1 := make([]int, 0)
	col2 := make([]int, 0)
	var val1 int
	var val2 int
	var err error
	col2Cnt := make(map[int]int, 0)
	for _, val := range lines {
		splits := strings.Split(val, "   ")
		if len(splits) != 2 {
			fmt.Println("err len")
			continue
		}
		val1, err = strconv.Atoi(splits[0])
		if err != nil {
			fmt.Println("err val1")
			continue
		}
		val2, err = strconv.Atoi(splits[1])
		if err != nil {
			fmt.Println("err val2")
			continue
		}
		col1 = append(col1, val1)
		col2 = append(col2, val2)
		col2Cnt[val2]++
	}
	sort.Ints(col1)
	sort.Ints(col2)

	if len(col1) != len(col2) {
		return "diff in length"
	}
	if isStep1 {
		var dist float64
		for i := range col1 {
			dist += math.Abs(float64(col1[i] - col2[i]))
		}
		return fmt.Sprintf("%d", int(dist))
	} else {
		var result int
		for _, val := range col1 {
			cnt, exist := col2Cnt[val]
			if exist {
				result += cnt * val
			}
		}
		return fmt.Sprintf("%d", result)
	}
}
