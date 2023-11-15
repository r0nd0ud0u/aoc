package main

import (
	"aoc2022/utils"
	"fmt"
)

func main() {
	fmt.Println(step1("03_test.txt"))
	fmt.Println(step1("03.txt"))
	fmt.Println(step2("03_test.txt"))
	fmt.Println(step2("03.txt"))
}

func step1(filename string) int {
	result := 0
	lines := utils.ReadInputLines(filename)
	for _, line := range lines {
		result += getFirstCommonChar(line[0:len(line)/2], line[len(line)/2:])
	}
	return result
}

func step2(filename string) int {
	result := 0
	lines := utils.ReadInputLines(filename)

	for i := 0; i < len(lines)-2; i = i + 3 {
		result1 := getAllCommonChar(lines[i], lines[i+1])
		result2 := getAllCommonChar(lines[i+1], lines[i+2])
		for _, i := range result1 {
			if utils.Contains[int](result2, i) {
				result += i
				break
			}
		}
	}

	return result
}

func getFirstCommonChar(str1 string, str2 string) int {
	result := 0
	for _, char1 := range str1 {
		found := false
		for _, char2 := range str2 {
			if char1 == char2 {
				found = true
				break
			}
		}
		if found {
			result += getPriority(byte(char1))
			break
		}
	}

	return result
}

func getAllCommonChar(str1 string, str2 string) []int {
	result := make([]int, 0)
	for _, char1 := range str1 {
		found := false
		for _, char2 := range str2 {
			if char1 == char2 {
				found = true
				break
			}
		}
		if found {
			result = append(result, getPriority(byte(char1)))
		}
	}

	return result
}

func getPriority(c byte) int {
	if c >= 65 && c <= 90 {
		return int(c - 64 + 26)
	}
	if c >= 97 && c <= 122 {
		return int(c - 96)
	}
	return 0
}
