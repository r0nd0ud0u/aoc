package utils

import (
	"fmt"
	"os"
	"strings"
)

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
