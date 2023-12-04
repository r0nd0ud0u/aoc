package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"gitlab.com/Alfred456654/aoc-utils"
)

func process(filename string, step1 bool) int {
	lines := aoc.ReadInputLines(filename)
	var result int = 0
	mapCopies := make(map[int]int, len(lines))

	re := regexp.MustCompile(`Card\s+(\d+): ([^\|]+) \| (.+)`)
	for i, line := range lines {
		m := re.FindStringSubmatch(line)
		if len(m) >= 4 {
			winArr := aoc.NewSet[int]()
			nbsArr := aoc.NewSet[int]()
			for _, w := range strings.Split(m[2], " ") {
				if len(w) == 0 {
					continue
				}
				wI, _ := strconv.Atoi(w)
				winArr.Add(wI)
			}
			for _, nb := range strings.Split(m[3], " ") {
				if len(nb) == 0 {
					continue
				}
				nbI, _ := strconv.Atoi(nb)
				nbsArr.Add(nbI)
			}
			inter := winArr.Intersection(nbsArr)
			if step1 {
				result += int(math.Pow(2., float64(inter.Size()-1)))
			} else {
				mapCopies[i]++
				for j := 0; j < inter.Size(); j++ {
					mapCopies[i+1+j] += mapCopies[i]
				}
			}
		}
	}

	if !step1 {
		for _, i := range mapCopies {
			result += i
		}
	}

	return result
}

func main() {
	aoc.CheckEquals(13, process("04_test1.txt", true))
	aoc.CheckEquals(30, process("04_test1.txt", false))
	fmt.Println(process("04.txt", true))
	fmt.Println(process("04.txt", false))
}
