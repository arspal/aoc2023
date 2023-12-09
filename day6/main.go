package main

import (
	"fmt"
	"github.com/arspal/aoc2023/utils"
	"math"
	"strconv"
)

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
	times = utils.ParseNumbers(scanner.Text())
	scanner.Scan()
	distances = utils.ParseNumbers(scanner.Text())

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
