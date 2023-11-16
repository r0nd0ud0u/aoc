package utils

import (
	"fmt"
	"log"
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

func ReadInputLinesWithoutTrim(filename string) []string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		Err("could not read file '%s'", err, filename)
	}
	allLines := string(bytes)
	linesRaw := strings.Split(allLines, "\n")
	return linesRaw
}

func Err(message string, err error, args ...interface{}) {
	logger.Printf(fmt.Sprintf("%s\n", message), args...)
	logger.Println(err)
}

var logger = log.New(os.Stderr, "", 0)

func Log(message string, args ...interface{}) {
	logger.Printf(fmt.Sprintf("%s\n", message), args...)
}
