package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/arspal/aoc2023/utils"
)

func main() {
	scanner := utils.OpenWithScanner("input.txt")
	history := make([][]int, 0)

	for scanner.Scan() {
		seq := make([]int, 0)

		for _, str := range strings.Split(scanner.Text(), " ") {
			v, err := strconv.Atoi(str)
			if err != nil {
				log.Panic("failed to parse input number", err)
			}
			seq = append(seq, v)
		}

		history = append(history, seq)
	}

	part1 := 0
	part2 := 0

	for _, seq := range history {
		part1 += nextValue(seq, false)
		part2 += nextValue(seq, true)
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func nextValue(seq []int, part2 bool) int {
	sequences := make([][]int, 0)
	currentSeq := seq

	for !allZeros(currentSeq) {
		sequences = append(sequences, currentSeq)

		nextSeq := make([]int, 0)
		for i, l := 0, len(currentSeq)-1; i < l; i += 1 {
			curr := currentSeq[i]
			next := currentSeq[i+1]
			nextSeq = append(nextSeq, next-curr)
		}

		currentSeq = nextSeq
	}

	value := 0
	for i := len(sequences) - 1; i >= 0; i -= 1 {
		if part2 {
			value = sequences[i][0] - value
		} else {
			value += sequences[i][len(sequences[i])-1]
		}
	}

	return value
}

func allZeros(seq []int) bool {
	for _, n := range seq {
		if n != 0 {
			return false
		}
	}

	return true
}
