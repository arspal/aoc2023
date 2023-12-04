package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
