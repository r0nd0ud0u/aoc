package main

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/Alfred456654/aoc-utils"
)

func main() {
	fmt.Println(step1("04_test.txt"))
	fmt.Println(step1("04.txt"))
	//fmt.Println(step2("04_test.txt"))
	//fmt.Println(step2("04.txt"))
}

func step1(filename string) int {
	result := 0

	lines := aoc.ReadInputLines(filename)
	for _, line := range lines {
		strs := strings.Split(line, ",")
		set1 := getAllInt(strs[0])
		set2 := getAllInt(strs[1])
		set3 := set1.Union(set2)
		if set3.Size() == set2.Size() && (set3.Size() == set1.Size() || set3.Size() == set2.Size()) {
			result++
		}
	}
	return result
}

func step2(filename string) int {
	result := 0

	return result
}

func getAllInt(str string) *aoc.Set[int] {
	result := aoc.NewSet[int]()
	strs := strings.Split(str, "-")
	if strs[0] == "" || strs[1] == "" {
		return result
	}
	index1, err1 := strconv.Atoi(strs[0])
	index2, err2 := strconv.Atoi(strs[1])

	if err1 == nil && err2 == nil {
		for i := index1; i < index2; i++ {
			result.Add(i)
		}
	} else {
		fmt.Println("oups")
	}
	return result
}
