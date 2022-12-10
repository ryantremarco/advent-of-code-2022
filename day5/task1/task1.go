package task1

import (
	"advent-of-code-2022/m/v2/day5/models"
	"advent-of-code-2022/m/v2/utils"
	"strings"
)

func TopStackItems(stacks []utils.List[string], moves []models.Move) string {
	for _, move := range moves {
		for i := 0; i < move.Count; i++ {
			popped := stacks[move.From-1].Pop()
			stacks[move.To-1].Push(popped)
		}
	}

	topItems := make([]string, len(stacks))
	for i := range stacks {
		topItems[i] = stacks[i].Pop()
	}

	return strings.Join(topItems, "")
}
