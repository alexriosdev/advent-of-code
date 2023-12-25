package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"

	"github.com/emirpasic/gods/queues/linkedlistqueue"
)

type module struct {
	symbol  rune
	name    string
	pulse   bool
	outputs []string
	inputs  map[string]bool
}

type input struct {
	source, dest string
	pulse        bool
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 20 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	replacer := strings.NewReplacer("->", "", ",", "")
	modules := map[string]module{}
	for _, line := range lines {
		config := strings.Fields(replacer.Replace(line))
		symbol, name := rune(config[0][0]), config[0][1:]
		if symbol == 'b' {
			name = "broadcaster"
		}
		module := module{
			symbol:  symbol,
			name:    name,
			pulse:   false,
			inputs:  map[string]bool{},
			outputs: config[1:],
		}
		modules[module.name] = module
	}
	for name, module := range modules {
		for _, dest := range module.outputs {
			if modules[dest].symbol == '&' {
				modules[dest].inputs[name] = false
			}
		}
	}
	low, high := 0, 0
	queue := linkedlistqueue.New()
	for i := 0; i < 1000; i++ {
		queue.Enqueue(input{
			source: "button",
			dest:   "broadcaster",
			pulse:  false,
		})
		for !queue.Empty() {
			val, _ := queue.Dequeue()
			curr := val.(input)
			if curr.pulse == false {
				low++
			} else {
				high++
			}
			if module, ok := modules[curr.dest]; ok {
				switch module.symbol {
				case '%':
					if curr.pulse == false {
						module.pulse = !module.pulse
						modules[curr.dest] = module
						for _, output := range module.outputs {
							queue.Enqueue(input{
								source: module.name,
								dest:   output,
								pulse:  module.pulse,
							})
						}
					}
				case '&':
					module.inputs[curr.source] = curr.pulse
					modules[curr.dest] = module
					for _, output := range module.outputs {
						queue.Enqueue(input{
							source: module.name,
							dest:   output,
							pulse:  !isAllTrue(module.inputs),
						})
					}
				default:
					for _, output := range module.outputs {
						queue.Enqueue(input{
							source: module.name,
							dest:   output,
							pulse:  curr.pulse,
						})
					}
				}
			}
		}
	}
	return low * high
}

func isAllTrue(inputs map[string]bool) bool {
	for _, v := range inputs {
		if !v {
			return false
		}
	}
	return true
}
