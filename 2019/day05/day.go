package day05

import (
	"github.com/ags131/adventofcode/2019/aoc"
	"github.com/ags131/adventofcode/2019/aoc/intcode"
)

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	program := intcode.ReadInput(input)

	part1 := 0
	part2 := 0
	m1 := intcode.NewMachine(program)
	m1.Input <- 1
	for v := range m1.Output {
		part1 = v
	}
	m2 := intcode.NewMachine(program)
	m2.Input <- 5
	for v := range m2.Output {
		part2 = v
	}
	return aoc.Output{Part1: part1, Part2: part2}
}
