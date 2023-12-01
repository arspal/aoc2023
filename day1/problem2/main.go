package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numbersAsText = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func toInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return num
}

func main() {
	file, err := os.OpenFile("../input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	sum := 0

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

		if indexF != 0 || indexL != length-1 {
			for i, num := range numbersAsText {
				idxF := strings.Index(text, num)
				idxL := strings.LastIndex(text, num)
				if idxF > -1 && idxF < indexF {
					numF = strconv.Itoa(i + 1)
					indexF = idxF
				}

				if idxL > -1 && idxL > indexL {
					numL = strconv.Itoa(i + 1)
					indexL = idxL
				}
			}
		}

		fmt.Printf("%s %s%s\n", text, numF, numL)
		sum += toInt(numF + numL)
	}
	fmt.Println()
	fmt.Println("total:", sum)
}
