package main

import (
	"advent-of-code-2022/m/v2/day5/models"
	"advent-of-code-2022/m/v2/day5/task1"
	"advent-of-code-2022/m/v2/day5/task2"
	"advent-of-code-2022/m/v2/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	stacks, moves, err := inputToStacksAndMoves(input)
	if err != nil {
		panic(err)
	}

	part2Stacks := make([]utils.List[string], len(stacks))
	for i := range stacks {
		part2Stacks[i] = utils.ListFrom(stacks[i].Slice())
	}
	topItems := task1.TopStackItems(stacks, moves)
	println("Top Stack Items", topItems)

	multimoveTopItems := task2.MultimoveTopStackItems(part2Stacks, moves)
	println("Multimove Top Stack Items", multimoveTopItems)
}

func inputToStacksAndMoves(input string) ([]utils.List[string], []models.Move, error) {
	split := strings.Split(input, "\n\n")
	if len(split) != 2 {
		return nil, nil, fmt.Errorf("expected stacks and moves input, got %#v", split)
	}

	stackInput := split[0]
	stackLines := strings.Split(stackInput, "\n")
	stackLines = stackLines[0 : len(stackLines)-1]
	stackCount := (len(stackLines[0]) + 1) / 4
	stacks := make([]utils.List[string], stackCount)

	moveInput := split[1]
	moveLines := strings.Split(moveInput, "\n")
	moves := make([]models.Move, len(moveLines))

	for i := len(stackLines) - 1; i >= 0; i-- {
		stackLineChars := strings.Split(stackLines[i], "")
		for i, j := 0, 1; i < len(stacks); i, j = i+1, j+4 {
			if stackLineChars[j] == " " {
				continue
			}
			stacks[i].Push(stackLineChars[j])
		}
	}

	for i, moveLine := range moveLines {
		words := strings.Split(moveLine, " ")
		count, err := strconv.Atoi(words[1])
		if err != nil {
			return nil, nil, err
		}
		from, err := strconv.Atoi(words[3])
		if err != nil {
			return nil, nil, err
		}
		to, err := strconv.Atoi(words[5])
		if err != nil {
			return nil, nil, err
		}

		moves[i] = models.Move{
			Count: count,
			From:  from,
			To:    to,
		}
	}

	return stacks, moves, nil
}
