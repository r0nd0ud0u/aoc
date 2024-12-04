package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type info struct {
	dest int64
	src  int64
	len  int64
}

func process(filename string) int64 {
	bytes, _ := os.ReadFile(filename)
	blocks := strings.Split(string(bytes), "\n\n")
	items := make(map[int][]string, len(blocks))

	instructions := make(map[int][]info, 0)

	for i, a := range blocks {
		linesRaw := strings.Split(a, "\n")
		items[i] = linesRaw
	}
	var res int64 = math.MaxInt64
	seeds := strings.Split(strings.Split(items[0][0], ":")[1], " ")

	for i := 1; i < len(items); i++ {
		arr := make([]info, 0)
		for _, elem := range items[i] {
			if strings.Contains(elem, ":") {
				continue
			}
			nbs := strings.Split(elem, " ")
			dest, err1 := strconv.ParseInt(nbs[0], 10, 64)
			src, err2 := strconv.ParseInt(nbs[1], 10, 64)
			length, err3 := strconv.ParseInt(nbs[2], 10, 64)
			if err1 != nil || err2 != nil || err3 != nil {
				fmt.Println("err decode")
			}
			inst := info{dest: dest, src: src, len: length}

			arr = append(arr, inst)
		}
		instructions[i] = arr
	}
	for _, seed := range seeds {
		nb, err := strconv.ParseInt(seed, 10, 64)
		if err != nil {
			continue
		}
		for i := 0; i < len(instructions); i++ {
			for _, elem := range instructions[i] {
				if nb >= elem.src && nb < elem.src+elem.len {
					nb = elem.dest - elem.src + nb
					break
				}
			}

		}

		if nb < res {
			res = nb
		}
	}

	return res
}

func main() {
	//aoc.CheckEquals(35, process("05_test1.txt"))
	//aoc.CheckEquals(30, process("04_test1.txt", false))
	fmt.Println(process("05.txt"))
	//fmt.Println(process("04.txt", false))
}
