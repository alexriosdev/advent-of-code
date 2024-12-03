package main

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type pair struct {
	low, high int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("2023 Day 19 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	parts := []map[rune]int{}
	ratingsReplacer := strings.NewReplacer("{x=", "", ",m=", " ", ",a=", " ", ",s=", " ", "}", "")
	for _, s := range strings.Split(lines[1], "\n") {
		ratings := strings.Fields(ratingsReplacer.Replace(s))
		parts = append(parts, map[rune]int{
			'x': utils.StrToInt(ratings[0]),
			'm': utils.StrToInt(ratings[1]),
			'a': utils.StrToInt(ratings[2]),
			's': utils.StrToInt(ratings[3]),
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

func part2(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	part := map[rune]pair{
		'x': {1, 4000},
		'm': {1, 4000},
		'a': {1, 4000},
		's': {1, 4000},
	}
	workflows := map[string][]string{}
	workflowsReplacer := strings.NewReplacer("{", " ", ",", " ", "}", "")
	for _, s := range strings.Split(lines[0], "\n") {
		split := strings.Fields(workflowsReplacer.Replace(s))
		workflows[split[0]] = split[1:]
	}
	return getCombinations(&part, &workflows, "in")
}

func isAccepted(part *map[rune]int, workflows *map[string][]string, rule string) bool {
	if rule == "A" {
		return true
	}
	if rule != "R" {
		for _, work := range (*workflows)[rule] {
			split := strings.Split(work, ":")
			if len(split) != 2 {
				return isAccepted(part, workflows, work)
			}
			key, op, num := rune(work[0]), work[1], runesToInt([]rune(work[2:]))
			if (op == '<' && (*part)[key] < num) || (op == '>' && (*part)[key] > num) {
				return isAccepted(part, workflows, split[1])
			}
		}
	}
	return false
}

func getCombinations(part *map[rune]pair, workflows *map[string][]string, rule string) int {
	if rule == "A" {
		prod := 1
		for _, r := range *part {
			prod *= r.high - r.low + 1
		}
		return prod
	}
	if rule != "R" {
		result := 0
		for _, work := range (*workflows)[rule] {
			split := strings.Split(work, ":")
			if len(split) != 2 {
				result += getCombinations(part, workflows, work)
				continue
			}
			key, op, num := rune(work[0]), work[1], runesToInt([]rune(work[2:]))
			r, t, f := (*part)[key], pair{}, pair{}
			if op == '<' {
				t.low, t.high = r.low, num-1
				f.low, f.high = num, r.high
			} else if op == '>' {
				t.low, t.high = num+1, r.high
				f.low, f.high = r.low, num
			}
			if t.low <= t.high {
				newPart := map[rune]pair{}
				for k, v := range *part {
					newPart[k] = v
				}
				newPart[key] = t
				result += getCombinations(&newPart, workflows, split[1])
			}
			if f.low <= f.high {
				(*part)[key] = f
			}
		}
		return result
	}
	return 0
}

func runesToInt(runes []rune) int {
	num := 0
	for i := 0; i < len(runes) && unicode.IsDigit(runes[i]); i++ {
		num = (num * 10) + int(runes[i]-'0')
	}
	return num
}
