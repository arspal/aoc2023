package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

func main() {
	file, _ := os.OpenFile("../input.txt", os.O_RDONLY, 0644)
	fileScanner := bufio.NewScanner(file)
	stringReader := strings.NewReader("")
	var tokensScanner scanner.Scanner

	setsPower := 0

	for fileScanner.Scan() {
		var maxRed, maxGreen, maxBlue int
		var red, green, blue int
		var amount int

		line := fileScanner.Text()
		stringReader.Reset(line)
		tokensScanner.Init(stringReader)
		tokensScanner.Whitespace ^= 1 << '\n'
		tokensScanner.Whitespace ^= 1<<',' | 1<<':'

		for {
			token := tokensScanner.Scan()
			text := tokensScanner.TokenText()

			if token == scanner.EOF {
				maxRed = max(maxRed, red)
				maxGreen = max(maxGreen, green)
				maxBlue = max(maxBlue, blue)
				power := maxRed * maxGreen * maxBlue

				setsPower += power
				break
			}

			switch token {
			case ';':
				maxRed = max(maxRed, red)
				maxGreen = max(maxGreen, green)
				maxBlue = max(maxBlue, blue)
			case scanner.Ident:
				if text == "red" {
					red = amount
				} else if text == "green" {
					green = amount
				} else if text == "blue" {
					blue = amount
				}
			case scanner.Int:
				val, _ := strconv.Atoi(text)
				amount = val
			default:
				fmt.Printf("Other: %s\n", text)
			}
		}
	}
	fmt.Println(setsPower)
}
