package utils

import (
	"fmt"
	"os"
	"strings"
)

type Applicable interface {
	int | int64 | float64 | string
}

func ReadInputLines(filename string) []string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Errorf("could not read file '%s'", err, filename)
	}
	allLines := string(bytes)
	linesRaw := strings.Split(allLines, "\n")
	lines := make([]string, 0)
	for _, line := range linesRaw {
		l := strings.Trim(line, " ")
		if len(l) > 0 {
			lines = append(lines, l)
		}
	}
	return lines
}

func Contains[T Applicable](values []T, value T) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
