package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/arspal/aoc2023/utils"
)

type Node struct {
	name  string
	left  string
	right string
}

func main() {
	scanner := utils.OpenWithScanner("input.txt")
	var instructions string
	var instructionsLength int
	var part1Node *Node
	part2Nodes := make([]*Node, 0, 5)
	var part1, part2 int
	nodesMap := make(map[string]*Node, 10)

	scanner.Scan()
	instructions = scanner.Text()
	instructionsLength = len(instructions)
	scanner.Scan()

	for scanner.Scan() {
		nodeText := scanner.Text()
		name, node := parseNode(nodeText)
		nodesMap[name] = node
		if name == "AAA" {
			part1Node = node
		}
		if strings.HasSuffix(name, "A") {
			part2Nodes = append(part2Nodes, node)
		}
	}
	{
		i := 0

		for {
			part1 += 1
			direction := instructions[i]
			switch direction {
			case 'R':
				part1Node = nodesMap[part1Node.right]
			case 'L':
				part1Node = nodesMap[part1Node.left]
			default:
				log.Panicln("encountered unknown direction:", direction)
			}

			if part1Node.name == "ZZZ" {
				break
			}

			if i+1 >= instructionsLength {
				i = 0
			} else {
				i += 1
			}
		}
	}

	steps := make([]int, 0, len(part2Nodes))
	for _, node := range part2Nodes {
		i := 0
		step := 0

		for {
			step += 1
			direction := instructions[i]
			switch direction {
			case 'R':
				node = nodesMap[node.right]
			case 'L':
				node = nodesMap[node.left]
			default:
				log.Panicln("encountered unknown direction:", direction)
			}

			if strings.HasSuffix(node.name, "Z") {
				break
			}

			if i+1 >= instructionsLength {
				i = 0
			} else {
				i += 1
			}
		}
		steps = append(steps, step)
	}

	part2 = lcm(steps...)

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func parseNode(str string) (string, *Node) {
	parts := strings.Split(str, " = ")
	nodeName := parts[0]
	nodeRefs := strings.Split(strings.Trim(parts[1], "()"), ", ")
	return nodeName, &Node{name: nodeName, left: nodeRefs[0], right: nodeRefs[1]}
}

func gcd(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

func lcmInner(a, b int) int {
	if a > b {
		return (a / gcd(a, b)) * b
	}
	return (b / gcd(a, b)) * a
}

func lcm(integers ...int) int {
	result := 1

	for i := 0; i < len(integers); i++ {
		result = lcmInner(result, integers[i])
	}

	return result
}
