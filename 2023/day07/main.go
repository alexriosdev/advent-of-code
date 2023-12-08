package main

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type play struct {
	hand          string
	handType, bid int
}

const (
	fiveOfAKind  = 7
	fourOfAKind  = 6
	fullHouse    = 5
	threeOfAKind = 4
	twoPair      = 3
	onePair      = 2
	highCard     = 1
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 07 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	plays := []play{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		plays = append(plays, play{
			hand: split[0],
			bid:  strToInt(split[1]),
		})
	}

	for i, play := range plays {
		freq := map[rune]int{}
		for _, c := range play.hand {
			freq[c]++
		}
		n := len(freq)
		switch {
		case n == 1:
			plays[i].handType = fiveOfAKind
		case n == 2 && containsValue(freq, 4):
			plays[i].handType = fourOfAKind
		case n == 2:
			plays[i].handType = fullHouse
		case n == 3 && containsValue(freq, 3):
			plays[i].handType = threeOfAKind
		case n == 3:
			plays[i].handType = twoPair
		case n == 4:
			plays[i].handType = onePair
		default:
			plays[i].handType = highCard
		}
	}

	cardMap := map[rune]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'J': 9,
		'T': 8,
		'9': 7,
		'8': 6,
		'7': 5,
		'6': 4,
		'5': 3,
		'4': 2,
		'3': 1,
		'2': 0,
	}
	sort.Slice(plays, func(a, b int) bool {
		return comparePlays(plays[a], plays[b], cardMap)
	})

	result := 0
	for i, play := range plays {
		result += play.bid * (i + 1)
	}
	return result
}

func comparePlays(playA, playB play, cardMap map[rune]int) bool {
	if playA.handType != playB.handType {
		return playA.handType < playB.handType
	}
	for i := range playA.hand {
		cardA, cardB := rune(playA.hand[i]), rune(playB.hand[i])
		if cardA != cardB {
			return cardMap[cardA] < cardMap[cardB]
		}
	}
	return false
}

func containsValue(freq map[rune]int, val int) bool {
	for _, v := range freq {
		if v == val {
			return true
		}
	}
	return false
}

func strToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
