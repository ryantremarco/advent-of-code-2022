package task1_test

import (
	"advent-of-code-2022/m/v2/day2/task1"
	"testing"
)

func TestTotalScore(t *testing.T) {
	exampleRounds := [][]string{
		{
			"A",
			"Y",
		},
		{
			"B",
			"X",
		},
		{
			"C",
			"Z",
		},
	}

	want := 15
	got := task1.TotalScore(exampleRounds)

	if want != got {
		t.Fatalf("expected score of %d, got %d", want, got)
	}
}
