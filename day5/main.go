package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"
	"unicode"

	"github.com/arspal/aoc2023/utils"
)

func parseIntegers(str string) []int {
	result := make([]int, 0)
	numbersAsText := strings.Split(str, " ")
	for _, num := range numbersAsText {
		result = append(result, utils.ParseInt(num))
	}

	return result
}

func main() {
	start := time.Now()
	scanner := utils.OpenWithScanner("input.test.txt")

	var seeds []int
	var seedsLength int
	mappings := make([][][3]int, 0)

	// parse seeds
	{
		scanner.Scan()
		seeds = parseIntegers(strings.Split(scanner.Text(), ": ")[1])
		seedsLength = len(seeds)
		if seedsLength%2 != 0 {
			log.Fatal("input seeds are unbalanced")
		}
	}

	// parse maps
	{
		var currentMap [][3]int

		scanner.Scan()
		scanner.Scan()
		for scanner.Scan() {
			text := scanner.Text()

			if text == "" {
				mappings = append(mappings, currentMap)
				currentMap = make([][3]int, 0)
				scanner.Scan()
				continue
			}

			var mapItems [3]int
			currentNumber := 0
			idx := 0

			for _, char := range text {
				if unicode.IsDigit(char) {
					currentNumber = currentNumber*10 + int(char-'0')
				} else {
					mapItems[idx] = currentNumber
					idx += 1
					currentNumber = 0
				}
			}

			mapItems[idx] = currentNumber

			currentMap = append(currentMap, mapItems)
		}

		mappings = append(mappings, currentMap)
	}

	locPart1 := int(math.Inf(1))
	locPart2 := int(math.Inf(1))

	for i, l := 0, seedsLength; i < l; i += 2 {
		locPart1 = min(locPart1, mapSeed(seeds[i], &mappings), mapSeed(seeds[i+1], &mappings))

		for seed, maxSeed := seeds[i], seeds[i]+seeds[i+1]; seed <= maxSeed; seed += 1 {
			locPart2 = min(locPart2, mapSeed(seed, &mappings))
		}
	}

	fmt.Println("part1:", locPart1)
	fmt.Println("part2:", locPart2)

	fmt.Println("execution time:", time.Since(start))
}

func mapSeed(seed int, mappings *[][][3]int) int {
	mapped := seed

	for _, mapping := range *mappings {
		for _, rules := range mapping {
			dest := rules[0]
			source := rules[1]
			amount := rules[2]

			if mapped >= source && mapped <= source+amount {
				mapped = mapped + (dest - source)
				break
			}
		}
	}
	return mapped
}
