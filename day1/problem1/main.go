package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

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

		for i := 0; i < length; i += 1 {
			char := text[i]
			if unicode.IsDigit(rune(char)) {
				numF = string(char)
				break
			}
		}
		for i := length - 1; i >= 0; i -= 1 {
			char := text[i]
			if unicode.IsDigit(rune(char)) {
				numL = string(char)
				break
			}
		}

		fmt.Printf("%s %s%s\n", text, numF, numL)

		num, err := strconv.Atoi(numF + numL)
		if err != nil {
			panic(err)
		}

		sum += num
	}
	fmt.Println()
	fmt.Println("total:", sum)
}
