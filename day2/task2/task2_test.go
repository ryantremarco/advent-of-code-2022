package task2_test

import (
	"advent-of-code-2022/m/v2/day2/task2"
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

	want := 12
	got := task2.GuidedTotalScore(exampleRounds)

	if want != got {
		t.Fatalf("expected score of %d, got %d", want, got)
	}
}
