package main

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
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
		case n == 2 && checkValue(freq, 4):
			plays[i].handType = fourOfAKind
		case n == 2:
			plays[i].handType = fullHouse
		case n == 3 && checkValue(freq, 3):
			plays[i].handType = threeOfAKind
		case n == 3:
			plays[i].handType = twoPair
		case n == 4:
			plays[i].handType = onePair
		default:
			plays[i].handType = highCard
		}
	}

	sort.Slice(plays, func(a, b int) bool {
		return comparePlays(plays[a], plays[b])
	})

	result := 0
	for i, play := range plays {
		result += play.bid * (i + 1)
	}
	return result
}

var cardMap = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

var fiveOfAKind = 7
var fourOfAKind = 6
var fullHouse = 5
var threeOfAKind = 4
var twoPair = 3
var onePair = 2
var highCard = 1

type play struct {
	hand          string
	handType, bid int
}

func comparePlays(playA, playB play) bool {
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

func checkValue(freq map[rune]int, val int) bool {
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
