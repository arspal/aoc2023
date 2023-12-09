package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func ParseInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("failed to parse number %s, err: %v\n", str, err)
	}

	return val
}

func Open(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("failed to open input file, error:", err)
	}
	return file
}

func OpenWithScanner(path string) *bufio.Scanner {
	file := Open(path)
	return bufio.NewScanner(file)
}

func ParseNumbers(str string) []int {
	result := make([]int, 0)
	var num *int
	for _, char := range str {
		if unicode.IsDigit(char) {
			if num == nil {
				num = new(int)
			}
			*num = (*num)*10 + int(char-'0')
			continue
		}

		if num != nil {
			result = append(result, *num)
			num = nil
		}
	}

	if num != nil {
		result = append(result, *num)
		num = nil
	}

	return result
}

func ParseNumber(str string) int {
	result := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			result = result*10 + int(char-'0')
		}
	}

	return result
}
