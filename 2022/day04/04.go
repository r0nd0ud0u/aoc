package main

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/Alfred456654/aoc-utils"
)

func main() {
	fmt.Println(step1and2("04_test.txt"))
	fmt.Println(step1and2("04.txt"))
}

func step1and2(filename string) (int, uint32) {
	result := 0
	var result2 uint32 = 0
	lines := aoc.ReadInputLines(filename)
	for _, line := range lines {
		strs := strings.Split(line, ",")
		set1 := getAllInt(strs[0])
		set2 := getAllInt(strs[1])
		set3 := set1.Intersection(set2)
		if set3.Size() > 0 && (set3.Size() == set1.Size() || set3.Size() == set2.Size()) {
			result++
		}
		if set3.Size() > 0 {
			result2++
		}
	}
	return result, result2
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
		for i := index1; i <= index2; i++ {
			result.Add(i)
		}
	}

	return result
}
