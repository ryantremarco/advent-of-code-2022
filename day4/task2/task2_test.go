package task2_test

import (
	"advent-of-code-2022/m/v2/day4/models"
	"advent-of-code-2022/m/v2/day4/task2"
	"testing"
)

func TestTotalScore(t *testing.T) {
	exampleJobs := []models.Job{
		{
			Start: 2,
			End:   4,
		},
		{
			Start: 6,
			End:   8,
		},
		{
			Start: 2,
			End:   3,
		},
		{
			Start: 4,
			End:   5,
		},
		{
			Start: 5,
			End:   7,
		},
		{
			Start: 7,
			End:   9,
		},
		{
			Start: 2,
			End:   8,
		},
		{
			Start: 3,
			End:   7,
		},
		{
			Start: 6,
			End:   6,
		},
		{
			Start: 4,
			End:   6,
		},
		{
			Start: 2,
			End:   6,
		},
		{
			Start: 4,
			End:   8,
		},
	}

	want := 2
	got := task2.OverlappingPairsCount(exampleJobs)

	if want != got {
		t.Fatalf("expected total of %d, got %d", want, got)
	}
}
