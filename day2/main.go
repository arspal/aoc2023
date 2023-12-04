package main

import (
	"fmt"
	"strings"
	"text/scanner"

	"github.com/arspal/aoc2023/utils"
)

func main() {
	fileScanner := utils.OpenWithScanner("input.txt")
	stringReader := strings.NewReader("")
	var tokensScanner scanner.Scanner

	gameRed, gameGreen, gameBlue := 12, 13, 14
	gameSum := 0
	setsPower := 0

	for fileScanner.Scan() {
		var gameId int
		var maxRed, maxGreen, maxBlue int
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
				setsPower += maxRed * maxGreen * maxBlue
				if maxRed <= gameRed && maxGreen <= gameGreen && maxBlue <= gameBlue {
					gameSum += gameId
				}
				break
			}

			switch token {
			case scanner.Ident:
				if text == "Game" {
					tokensScanner.Scan()
					gameId = utils.ParseInt(tokensScanner.TokenText())
				} else if text == "red" {
					maxRed = max(maxRed, amount)
				} else if text == "green" {
					maxGreen = max(maxGreen, amount)
				} else if text == "blue" {
					maxBlue = max(maxBlue, amount)
				} else {
					fmt.Println("Unknown identifier: ", text)
				}
			case scanner.Int:
				amount = utils.ParseInt(text)
			default:
				// noop
			}
		}
	}

	fmt.Println("part1:", gameSum)
	fmt.Println("part2:", setsPower)
}
