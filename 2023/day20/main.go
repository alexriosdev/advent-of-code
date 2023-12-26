package main

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
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
	fmt.Printf("Part 2: %v\n", part2(lines))
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

func part2(lines []string) int {
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
	var feed string
	for name, module := range modules {
		if slices.Contains(module.outputs, "rx") {
			feed = name
			break
		}
	}
	visited := map[string]bool{}
	for name, module := range modules {
		if slices.Contains(module.outputs, feed) {
			visited[name] = false
		}
	}
	cycles := map[string]int{}
	queue := linkedlistqueue.New()
	for i := 1; ; i++ {
		queue.Enqueue(input{
			source: "button",
			dest:   "broadcaster",
			pulse:  false,
		})
		for !queue.Empty() {
			val, _ := queue.Dequeue()
			curr := val.(input)
			if module, ok := modules[curr.dest]; ok {
				if module.name == feed && curr.pulse == true {
					visited[curr.source] = true
					if _, ok := cycles[curr.source]; !ok {
						cycles[curr.source] = i
					}
					if isAllTrue(visited) {
						x := 1
						for _, v := range cycles {
							x = x * v / gcd(x, v)
						}
						return x
					}
				}
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
	return -1
}

func isAllTrue(m map[string]bool) bool {
	for _, v := range m {
		if !v {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
