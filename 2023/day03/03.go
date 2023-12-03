package main

import (
	"aoc/utils"
	"fmt"
	"strconv"

	"gitlab.com/Alfred456654/aoc-utils"
)

type partNumberPart2 struct {
	value int
	cnt   int
}

func process1(nbStr string, isInResult bool) int64 {
	nb, err := strconv.Atoi(nbStr)
	var res int64 = 0
	if err == nil && isInResult {
		res = int64(nb)
	}
	return res
}

func process(filename string, step1 bool) int64 {
	lines := aoc.ReadInputLines(filename)
	var result int64 = 0

	width := len(lines[0])
	height := len(lines)
	arr := make([]int, width*height)
	for i := range arr {
		arr[i] = 0
	}

	for i, line := range lines {
		for j, c := range line {
			if (!step1 && c == '*' && !utils.IsDigit(byte(c))) ||
				(step1 && c != '.' && !utils.IsDigit(byte(c))) {
				arr[j+i*len(line)] = 1
			}
		}
	}

	resultMap := make(map[int]partNumberPart2, 0)

	for i, line := range lines {
		nbStr := ""
		isInResult := false
		indexMap := -1
		for j, c := range line {

			if utils.IsDigit(byte(c)) {
				nbStr += string(c)
				if !isInResult {
					isInResult, indexMap = isPartNumber(j, i, arr, width, height)
				}
			} else {
				if step1 {
					result += process1(nbStr, isInResult)
				} else {
					process2(nbStr, isInResult, resultMap, indexMap)
				}
				nbStr = ""
				isInResult = false
				indexMap = -1
			}
		}
		if len(nbStr) > 0 {
			if step1 {
				result += process1(nbStr, isInResult)
			} else {
				process2(nbStr, isInResult, resultMap, indexMap)
			}
		}
	}

	for _, val := range resultMap {
		if val.cnt == 2 {
			result += int64(val.value)
		}
	}

	return result
}

func process2(nbStr string, isInResult bool, resultMap map[int]partNumberPart2, indexMap int) {
	nb, err := strconv.Atoi(nbStr)
	if err != nil {
		nb = 0
	} else if isInResult {
		val := 0
		if resultMap[indexMap].value == 0 {
			val = nb
		} else {
			val = resultMap[indexMap].value * nb
		}
		resultMap[indexMap] = partNumberPart2{
			value: val,
			cnt:   resultMap[indexMap].cnt + 1,
		}
	}
}

func main() {
	aoc.CheckEquals(4361, process("03_test1.txt", true))
	aoc.CheckEquals(467835, process("03_test1.txt", false))
	fmt.Println(process("03.txt", true))
	fmt.Println(process("03.txt", false))
}

func isPartNumber(i int, j int, arr []int, len int, h int) (bool, int) {
	res := false
	ind := 0
	if i+j*len+1 < (j+1)*len && arr[i+j*len+1] == 1 { //
		res = true
		ind = i + j*len + 1
	}
	if i+j*len-1 >= 0 && arr[i+j*len-1] == 1 { //
		res = true
		ind = i + j*len - 1
	}
	if i+(j-1)*len-1 >= 0 && i+(j-1)*len-1 <= j*len && arr[i+(j-1)*len-1] == 1 { //
		res = true
		ind = i + (j-1)*len - 1
	}
	if i+(j-1)*len >= 0 && arr[i+(j-1)*len] == 1 { //
		res = true
		ind = i + (j-1)*len
	}
	if i+(j-1)*len+1 >= 0 && i+(j-1)*len+1 < j*len && arr[i+(j-1)*len+1] == 1 { //
		res = true
		ind = i + (j-1)*len + 1
	}
	if i+(j+1)*len-1 >= (j+1)*len && i+(j+1)*len-1 < len*h && arr[i+(j+1)*len-1] == 1 { //
		res = true
		ind = i + (j+1)*len - 1
	}
	if i+(j+1)*len < len*h && arr[i+(j+1)*len] == 1 { //
		res = true
		ind = i + (j+1)*len
	}
	if i+(j+1)*len+1 < len*h && i+(j+1)*len+1 < (j+2)*len && arr[i+(j+1)*len+1] == 1 { //
		res = true
		ind = i + (j+1)*len + 1
	}
	return res, ind
}
