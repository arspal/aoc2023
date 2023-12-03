package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/arspal/aoc2023/utils"
)

var numbersAsText = [10]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func main() {
	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("failed to open input file, error:", err)
	}

	scanner := bufio.NewScanner(file)
	sumP1 := 0
	sumP2 := 0

	for scanner.Scan() {
		text := scanner.Text()
		length := len(text)

		var numF, numL string
		var indexF, indexL int

		for i := 0; i < length; i += 1 {
			char := text[i]
			if unicode.IsDigit(rune(char)) {
				numF = string(char)
				indexF = i
				break
			}
		}
		for i := length - 1; i >= 0; i -= 1 {
			char := text[i]
			if unicode.IsDigit(rune(char)) {
				numL = string(char)
				indexL = i
				break
			}
		}

		if numF != "" && numL != "" {
			sumP1 += utils.ParseInt(numF + numL)
		}

		if indexF != 0 || indexL != length-1 {
			for i, num := range numbersAsText {
				idxF := strings.Index(text, num)
				idxL := strings.LastIndex(text, num)
				if idxF > -1 && idxF <= indexF {
					numF = strconv.Itoa(i + 1)
					indexF = idxF
				}

				if idxL > -1 && idxL >= indexL {
					numL = strconv.Itoa(i + 1)
					indexL = idxL
				}
			}
		}

		sumP2 += utils.ParseInt(numF + numL)
	}

	fmt.Println("part1:", sumP1)
	fmt.Println("part2:", sumP2)
}
