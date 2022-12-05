package task2

import (
	"advent-of-code-2022/m/v2/day5/models"
	"advent-of-code-2022/m/v2/utils"
	"strings"
)

func MultimoveTopStackItems(stacks []utils.Stack[string], moves []models.Move) string {
	for _, move := range moves {
		popped := stacks[move.From-1].PopN(move.Count)
		stacks[move.To-1].Push(popped...)
	}

	topItems := make([]string, len(stacks))
	for i := range stacks {
		topItems[i] = stacks[i].Pop()
	}

	return strings.Join(topItems, "")
}
