package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/arspal/aoc2023/utils"
)

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("failed to open input file, error:", err)
	}

	lines := strings.Split(string(contents), "\n")
	var total, twoPartTotal int

	for lineIdx, line := range lines {
		for charIdx, char := range line {
			if char == '.' || unicode.IsDigit(char) {
				continue
			}

			numbers := make([]int, 0, 8)

			parseAdjacentNumbers(charIdx, line, &numbers)
			if lineIdx > 0 {
				parseAdjacentNumbers(charIdx, lines[lineIdx-1], &numbers)
			}
			if lineIdx+1 <= len(lines) {
				parseAdjacentNumbers(charIdx, lines[lineIdx+1], &numbers)
			}

			for _, num := range numbers {
				total += num
			}

			if char == '*' && len(numbers) == 2 {
				twoPartTotal += numbers[0] * numbers[1]
			}
		}
	}

	fmt.Println("part1:", total)
	fmt.Println("part2:", twoPartTotal)
}

func parseAdjacentNumbers(middle int, line string, numbers *[]int) {
	left := middle - 1
	right := middle + 1

	if unicode.IsDigit(rune(line[middle])) {
		start, end := getNumberStartEnd(line, middle)
		*numbers = append(*numbers, utils.ParseInt(line[start:end]))
		return
	}

	if unicode.IsDigit(rune(line[left])) {
		start, end := getNumberStartEnd(line, left)
		*numbers = append(*numbers, utils.ParseInt(line[start:end]))
	}
	if unicode.IsDigit(rune(line[right])) {
		start, end := getNumberStartEnd(line, right)
		*numbers = append(*numbers, utils.ParseInt(line[start:end]))
	}
}

func getNumberStartEnd(line string, index int) (int, int) {
	start := index
	end := index + 1

	for i := start - 1; i >= 0 && unicode.IsDigit(rune(line[i])); i -= 1 {
		start = i
	}

	for end < len(line) && unicode.IsDigit(rune(line[end])) {
		end += 1
	}

	return start, end
}
