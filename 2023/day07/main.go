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

type handType struct {
	value int
	shape []int
}

var fiveOfAKind  = handType{value: 7, shape: []int{5}}
var fourOfAKind  = handType{value: 6, shape: []int{1, 4}}
var fullHouse    = handType{value: 5, shape: []int{2, 3}}
var threeOfAKind = handType{value: 4, shape: []int{1, 1, 3}}
var twoPair      = handType{value: 3, shape: []int{1, 2, 2}}
var onePair      = handType{value: 2, shape: []int{1, 1, 1, 2}}
var highCard     = handType{value: 1, shape: []int{1, 1, 1, 1, 1}}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 07 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
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
		counts := []int{}
		for _, v := range freq {
			counts = append(counts, v)
		}
		sort.Slice(counts, func(i, j int) bool {
			return counts[i] < counts[j]
		})
		switch {
		case sliceEqual(counts, fiveOfAKind.shape):
			plays[i].handType = fiveOfAKind.value
		case sliceEqual(counts, fourOfAKind.shape):
			plays[i].handType = fourOfAKind.value
		case sliceEqual(counts, fullHouse.shape):
			plays[i].handType = fullHouse.value
		case sliceEqual(counts, threeOfAKind.shape):
			plays[i].handType = threeOfAKind.value
		case sliceEqual(counts, twoPair.shape):
			plays[i].handType = twoPair.value
		case sliceEqual(counts, onePair.shape):
			plays[i].handType = onePair.value
		default:
			plays[i].handType = highCard.value
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

func part2(lines []string) int {
	plays := []play{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		plays = append(plays, play{
			hand: split[0],
			bid:  strToInt(split[1]),
		})
	}

	for i, play := range plays {
		freqA := map[rune]int{}
		freqB := map[rune]int{}
		for _, c := range play.hand {
			freqA[c]++
			freqB[c]++
		}
		target := 0
		for k, v := range freqA {
			if k != 'J' {
				target = getMax(target, v)
			}
		}
		if jCount, ok := freqA['J']; ok {
			for k, v := range freqA {
				if v == target && k != 'J' {
					freqB[k] += jCount
					delete(freqB, 'J')
					break
				}
			}
		}
		counts := []int{}
		for _, v := range freqB {
			counts = append(counts, v)
		}
		sort.Slice(counts, func(i, j int) bool {
			return counts[i] < counts[j]
		})
		switch {
		case sliceEqual(counts, fiveOfAKind.shape):
			plays[i].handType = fiveOfAKind.value
		case sliceEqual(counts, fourOfAKind.shape):
			plays[i].handType = fourOfAKind.value
		case sliceEqual(counts, fullHouse.shape):
			plays[i].handType = fullHouse.value
		case sliceEqual(counts, threeOfAKind.shape):
			plays[i].handType = threeOfAKind.value
		case sliceEqual(counts, twoPair.shape):
			plays[i].handType = twoPair.value
		case sliceEqual(counts, onePair.shape):
			plays[i].handType = onePair.value
		default:
			plays[i].handType = highCard.value
		}
	}

	cardMap := map[rune]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
		'J': 0,
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

func strToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
