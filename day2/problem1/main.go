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

	maxRed, maxGreen, maxBlue := 12, 13, 14
	possibleGames := make([]int, 0, 8)
	idxSum := 0

fileScan:
	for fileScanner.Scan() {
		var gameId int
		var red, green, blue int
		var isGame bool
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
				if red > maxRed || green > maxGreen || blue > maxBlue {
					continue fileScan
				}
				possibleGames = append(possibleGames, gameId)
				idxSum += gameId
				break
			}

			switch token {
			case ';':
				if red > maxRed || green > maxGreen || blue > maxBlue {
					continue fileScan
				} else {
					red, green, blue = 0, 0, 0
				}
			case scanner.Ident:
				if text == "Game" {
					isGame = true
				} else if text == "red" {
					red = amount
				} else if text == "green" {
					green = amount
				} else if text == "blue" {
					blue = amount
				} else {
					fmt.Println("Unknown identifier: ", text)
				}
			case scanner.Int:
				val, _ := strconv.Atoi(text)
				if isGame {
					gameId = val
					isGame = false
				} else {
					amount = val
				}
			default:
				// noop
			}
		}
	}

	fmt.Println(idxSum)
}
