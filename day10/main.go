package main

import (
	"advent-of-code-2022/m/v2/utils"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	instructions := inputToInstructions(input)
	println("Sum of Signal Strengths", sumOfStrengthsAtCycles(instructions, []int{20, 60, 100, 140, 180, 220}))

	drawCRT(instructions)
}

// part 1
func sumOfStrengthsAtCycles(instructions []Instruction, cyclesToSum []int) int {
	sum := 0
	cpu := CPU{
		Register: 1,
	}
	cpu.Load(instructions...)
	cycle := 0
	for cpu.IsRunning() {
		cycle++
		if utils.Contains(cyclesToSum, (cycle)) {
			sum += cpu.Register * (cycle)
		}
		cpu.Cycle()
	}
	return sum
}

// part 2
func drawCRT(instructions []Instruction) {
	cpu := CPU{
		Register: 1,
	}
	cpu.Load(instructions...)
	cycle := 0
	for cpu.IsRunning() {
		pixel := cycle % 40
		if pixel == 0 {
			println()
		}
		if utils.Contains([]int{cpu.Register - 1, cpu.Register, cpu.Register + 1}, pixel) {
			print("#")
		} else {
			print(".")
		}

		cpu.Cycle()
		cycle++
	}
}

type CPU struct {
	Cycling                utils.List[Instruction]
	CurrentInstrucitonTime int
	Register               int
}

func (c *CPU) Load(instructions ...Instruction) {
	c.Cycling = utils.ListFrom(instructions)
}

func (c *CPU) Cycle() {
	if c.Cycling.Length() > 0 {
		c.CurrentInstrucitonTime++
		current, err := c.Cycling.Get(0)
		if err != nil {
			panic("can't process missing instruction")
		}

		if c.CurrentInstrucitonTime >= current.Command.TimeToComplete {
			completed := c.Cycling.Dequeue()
			if completed.Command == Add {
				c.Register += completed.Argument
			}
			c.CurrentInstrucitonTime = 0
		}
	}
}

func (c *CPU) IsRunning() bool {
	if c.Cycling.Count() > 0 {
		return true
	}
	return false
}

type Command struct {
	TimeToComplete int
}

var (
	NoOp       Command = Command{1}
	Add        Command = Command{2}
	CommandMap         = map[string]Command{
		"noop": NoOp,
		"addx": Add,
	}
)

type Instruction struct {
	Command  Command
	Argument int
}

func inputToInstructions(input string) []Instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		command := CommandMap[parts[0]]
		argument := 0
		if len(parts) > 1 {
			var err error
			argument, err = strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
		}
		instructions[i] = Instruction{command, argument}
	}

	return instructions
}