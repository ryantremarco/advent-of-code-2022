package main

import (
	"advent-of-code-2022/m/v2/day2/task1"
	"advent-of-code-2022/m/v2/day2/task2"
	"strings"

	// "advent-of-code-2022/m/v2/day1/task2"
	"advent-of-code-2022/m/v2/utils"
	"fmt"
)

func main() {
	input, err := utils.ReadInputFile()
	if err != nil {
		panic(err)
	}

	rounds, err := inputToRounds(input)
	if err != nil {
		panic(err)
	}

	totalScore := task1.TotalScore(rounds)
	guidedTotalScore := task2.GuidedTotalScore(rounds)

	fmt.Println("Total Score", totalScore)
	fmt.Println("Guided Total Score", guidedTotalScore)
}

func inputToRounds(input []string) ([][]string, error) {
	rounds := make([][]string, len(input))
	for i, line := range input {
		moves := strings.Split(line, " ")
		if len(moves) != 2 {
			return nil, fmt.Errorf("Expected 2 moves per input line, but saw '%#v'", moves)
		}
		rounds[i] = moves
	}
	return rounds, nil
}
