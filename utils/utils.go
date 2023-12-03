package utils

import (
	"log"
	"strconv"
)

func ParseInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("failed to parse number %s, err: %v\n", str, err)
	}

	return val
}
