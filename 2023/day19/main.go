package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("2023 Day 19 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
}

func part1(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	parts := []map[rune]int{}
	ratingsReplacer := strings.NewReplacer("{x=", "", ",m=", " ", ",a=", " ", ",s=", " ", "}", "")
	for _, s := range strings.Split(lines[1], "\n") {
		ratings := strings.Fields(ratingsReplacer.Replace(s))
		parts = append(parts, map[rune]int{
			'x': strToInt(ratings[0]),
			'm': strToInt(ratings[1]),
			'a': strToInt(ratings[2]),
			's': strToInt(ratings[3]),
		})
	}
	workflows := map[string][]string{}
	workflowsReplacer := strings.NewReplacer("{", " ", ",", " ", "}", "")
	for _, s := range strings.Split(lines[0], "\n") {
		split := strings.Fields(workflowsReplacer.Replace(s))
		workflows[split[0]] = split[1:]
	}
	result := 0
	for _, part := range parts {
		if isAccepted(&part, &workflows, "in") {
			for _, v := range part {
				result += v
			}
		}
	}
	return result
}

func isAccepted(part *map[rune]int, workflows *map[string][]string, rule string) bool {
	if rule == "A" {
		return true
	}
	if rule != "R" {
		for _, work := range (*workflows)[rule] {
			split := strings.Split(work, ":")
			if len(split) == 2 && work[1] == '<' {
				if (*part)[rune(work[0])] < runesToInt([]rune(work[2:])) {
					return isAccepted(part, workflows, split[1])
				} else {
					continue
				}
			} else if len(split) == 2 && work[1] == '>' {
				if (*part)[rune(work[0])] > runesToInt([]rune(work[2:])) {
					return isAccepted(part, workflows, split[1])
				} else {
					continue
				}
			} else {
				return isAccepted(part, workflows, work)
			}
		}
	}
	return false
}

func runesToInt(runes []rune) int {
	num := 0
	for i := 0; i < len(runes) && unicode.IsDigit(runes[i]); i++ {
		num = (num * 10) + int(runes[i]-'0')
	}
	return num
}

func strToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
