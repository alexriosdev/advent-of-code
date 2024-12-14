package main

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var PRIZE_DELTA = 10000000000000

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("2024 Day 13 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []byte) int {
	games := getGames(input)
	return getTokens(games, 0)
}

func part2(input []byte) int {
	games := getGames(input)
	return getTokens(games, PRIZE_DELTA)
}

type coordinate struct {
	x, y int
}

type game struct {
	buttonA, buttonB, prize coordinate
}

func getGames(input []byte) []game {
	pattern := regexp.MustCompile(`X[\+=]\d+,\sY[\+=]\d+`)
	buttonReplacer := strings.NewReplacer("X+", "", ",", " ", "Y+", "")
	prizeReplacer := strings.NewReplacer("X=", "", ",", " ", "Y=", "")
	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	games := []game{}
	for _, section := range sections {
		lines := strings.Split(section, "\n")
		buttonA := strings.Fields(buttonReplacer.Replace(pattern.FindString(lines[0])))
		buttonB := strings.Fields(buttonReplacer.Replace(pattern.FindString(lines[1])))
		prize := strings.Fields(prizeReplacer.Replace(pattern.FindString(lines[2])))
		game := game{
			buttonA: coordinate{utils.StrToInt(buttonA[0]), utils.StrToInt(buttonA[1])},
			buttonB: coordinate{utils.StrToInt(buttonB[0]), utils.StrToInt(buttonB[1])},
			prize:   coordinate{utils.StrToInt(prize[0]), utils.StrToInt(prize[1])},
		}
		games = append(games, game)
	}
	return games
}

func getTokens(games []game, prizeDelta int) int {
	result := 0
	for _, game := range games {
		game.prize.x += prizeDelta
		game.prize.y += prizeDelta
		den := (game.buttonA.x * game.buttonB.y) - (game.buttonA.y * game.buttonB.x)
		numA := (game.buttonB.y * game.prize.x) - (game.buttonB.x * game.prize.y)
		numB := (game.buttonA.x * game.prize.y) - (game.buttonA.y * game.prize.x)
		if den == 0 || numA%den != 0 || numB%den != 0 {
			continue
		}
		a, b := numA/den, numB/den
		if prizeDelta != 0 || max(a, b) <= 100 {
			result += 3*a + b
		}
	}
	return result
}
