package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/arspal/aoc2023/utils"
)

func main() {
	scanner := utils.OpenWithScanner("input.txt")

	score := 0
	cardsTotal := 0

	cardNumber := 0
	cardIdToAmount := make(map[int]int)

	for scanner.Scan() {
		cardNumber += 1
		text := scanner.Text()

		cardIdToAmount[cardNumber] += 1
		cardNumbers := strings.Split(strings.Split(text, ": ")[1], " | ")
		winning := strings.Split(cardNumbers[0], " ")
		hand := strings.Split(cardNumbers[1], " ")

		cardScore := 0
		matches := 0

		for _, num := range hand {
			if num != "" && slices.Contains(winning, num) {
				matches += 1
				cardScore = max(1, cardScore*2)
			}
		}

		for i := 1; i <= matches; i += 1 {
			cardIdToAmount[cardNumber+i] += cardIdToAmount[cardNumber]
		}

		cardsTotal += cardIdToAmount[cardNumber]
		score += cardScore
	}

	fmt.Println("part1:", score)
	fmt.Println("part2:", cardsTotal)
}
