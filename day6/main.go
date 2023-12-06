package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"

	"github.com/arspal/aoc2023/utils"
)

func parseNumbers(str string) []int {
	result := make([]int, 0)
	var num *int
	for _, char := range str {
		if unicode.IsDigit(char) {
			if num == nil {
				num = new(int)
			}
			*num = (*num)*10 + int(char-'0')
			continue
		}

		if num != nil {
			result = append(result, *num)
			num = nil
		}
	}

	if num != nil {
		result = append(result, *num)
		num = nil
	}

	return result
}

func getWaysToWin(time, distance int) int {
	result := 0

	for t := 1; t <= time; t += 1 {
		if (time-t)*t > distance {
			result += 1
		}
	}

	return result
}

func main() {
	scanner := utils.OpenWithScanner("input.txt")

	var times, distances []int

	scanner.Scan()
	times = parseNumbers(scanner.Text())
	scanner.Scan()
	distances = parseNumbers(scanner.Text())

	p1result := 0
	p2result := 0

	{
		p1result = getWaysToWin(times[0], distances[0])
		for i, l := 1, len(times); i < l; i += 1 {
			time := times[i]
			distance := distances[i]
			ways := getWaysToWin(time, distance)
			p1result *= ways
		}
	}
	{
		time := 0
		distance := 0
		for i, l := 0, len(times); i < l; i += 1 {
			t := len(strconv.Itoa(times[i]))
			time = time*int(math.Pow(10, float64(t))) + times[i]
			d := len(strconv.Itoa(distances[i]))
			distance = distance*int(math.Pow(10, float64(d))) + distances[i]
		}

		p2result = getWaysToWin(time, distance)
	}

	fmt.Println("part1:", p1result)
	fmt.Println("part2:", p2result)
}
