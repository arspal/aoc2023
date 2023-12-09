package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/arspal/aoc2023/utils"
)

type Kind int

const (
	HandSize = 5
	Faces1   = "23456789TJQKA"
	Faces2   = "J23456789TQKA"

	KindFive      Kind = 1 << 7
	KindFour           = 1 << 6
	KindFullHouse      = 1 << 5
	KindThree          = 1 << 4
	KindTwoPair        = 1 << 3
	KindOnePair        = 1 << 2
	KindHighCard       = 1
)

type Hand struct {
	kind  Kind
	cards string
	bid   int
}

type CardGroup struct {
	singles  int
	pairs    int
	multiple int
	jokers   int
}

type FaceCount struct {
	face  rune
	count int
}

func main() {
	scanner := utils.OpenWithScanner("input.txt")
	hands := make([]Hand, 0, 4)

	counter := newFaceCounter()
	for scanner.Scan() {
		hand := parseHand(scanner.Text())
		counter.countFaces(hand, false)
		applyKind(hand, counter)
		hands = append(hands, *hand)
	}

	part1 := 0
	part2 := 0

	sort(hands, Faces1)

	for index, hand := range hands {
		part1 += (index + 1) * hand.bid
	}

	for index := range hands {
		counter.countFaces(&hands[index], true)
		applyKind(&hands[index], counter)
	}

	sort(hands, Faces2)

	for index, hand := range hands {
		part2 += (index + 1) * hand.bid
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

type FaceCounter struct {
	faces map[rune]*FaceCount
	group *CardGroup
}

func (counter *FaceCounter) countFaces(hand *Hand, countJokers bool) {
	clear(counter.faces)
	counter.group.multiple = 0
	counter.group.pairs = 0
	counter.group.singles = 0
	counter.group.jokers = 0

	for _, card := range hand.cards {
		if countJokers && card == 'J' {
			counter.group.jokers += 1
			continue
		}

		v, ok := counter.faces[card]
		if !ok {
			v = &FaceCount{face: card}
			counter.faces[card] = v
		}

		v.count += 1
	}
}

func newFaceCounter() *FaceCounter {
	return &FaceCounter{faces: make(map[rune]*FaceCount), group: &CardGroup{}}
}
func parseHand(str string) *Hand {
	parts := strings.Split(str, " ")
	hand := &Hand{cards: parts[0], bid: utils.ParseNumber(parts[1])}
	return hand
}

func (counter *FaceCounter) values() *[]FaceCount {
	result := make([]FaceCount, 0, len(counter.faces))

	for _, occur := range counter.faces {
		result = append(result, *occur)
	}
	return &result
}

func sort(hands []Hand, faces string) {
	slices.SortFunc(hands, func(a Hand, b Hand) int {
		if a.kind != b.kind {
			return int(a.kind - b.kind)
		}

		for i := 0; i < HandSize; i += 1 {
			idx1 := strings.Index(faces, string(a.cards[i]))
			idx2 := strings.Index(faces, string(b.cards[i]))
			if idx1 != idx2 {
				return idx1 - idx2
			}
		}

		return 0
	})
}

func applyKind(hand *Hand, counter *FaceCounter) {
	group := counter.group
	faces := counter.values()

	slices.SortFunc(*faces, func(a, b FaceCount) int {
		return b.count - a.count
	})

	for _, face := range *faces {
		if face.count > 2 {
			group.multiple = face.count
		} else if face.count > 1 {
			group.pairs += 1
		} else {
			group.singles += 1
		}
	}

	if group.multiple > 0 {
		if group.multiple+group.jokers == 5 {
			hand.kind = KindFive
		} else if group.multiple+group.jokers == 4 {
			hand.kind = KindFour
		} else if group.multiple+group.jokers == 3 {
			if group.pairs == 1 {
				hand.kind = KindFullHouse
			} else {
				hand.kind = KindThree
			}
		}
	} else if group.pairs > 0 {
		if group.pairs == 2 {
			if group.jokers > 0 {
				hand.kind = KindFullHouse
			} else {
				hand.kind = KindTwoPair
			}
		} else if group.jokers == 3 {
			hand.kind = KindFive
		} else if group.jokers == 2 {
			hand.kind = KindFour
		} else if group.jokers == 1 {
			hand.kind = KindThree
		} else {
			hand.kind = KindOnePair
		}
	} else if group.jokers > 0 {
		switch group.jokers {
		case 5:
			hand.kind = KindFive
		case 4:
			hand.kind = KindFive
		case 3:
			hand.kind = KindFour
		case 2:
			hand.kind = KindThree
		case 1:
			hand.kind = KindOnePair
		}
	} else {
		hand.kind = KindHighCard
	}
}
