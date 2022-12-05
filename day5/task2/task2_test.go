package task2_test

import (
	"advent-of-code-2022/m/v2/day5/models"
	"advent-of-code-2022/m/v2/day5/task2"
	"advent-of-code-2022/m/v2/utils"
	"testing"
)

func TestTopStackItems(t *testing.T) {
	exampleStacks := []utils.Stack[string]{
		utils.StackFrom([]string{"Z", "N"}),
		utils.StackFrom([]string{"M", "C", "D"}),
		utils.StackFrom([]string{"P"}),
	}
	exampleMoves := []models.Move{
		{
			Count: 1,
			From:  2,
			To:    1,
		},
		{
			Count: 3,
			From:  1,
			To:    3,
		},
		{
			Count: 2,
			From:  2,
			To:    1,
		},
		{
			Count: 1,
			From:  1,
			To:    2,
		},
	}
	want := "MCD"
	got := task2.MultimoveTopStackItems(exampleStacks, exampleMoves)

	if want != got {
		t.Fatalf("expected top items of %s, got %s", want, got)
	}
}
