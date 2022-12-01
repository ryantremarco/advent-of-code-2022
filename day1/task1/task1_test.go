package task1_test

import (
	"advent-of-code-2022/m/v2/day1/models"
	"advent-of-code-2022/m/v2/day1/task1"
	"testing"
)

func TestMaxCalories(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		exampleElves := []models.Elf{
			{
				FoodCalories: []int{1000, 2000, 3000},
			},
			{
				FoodCalories: []int{4000},
			},
			{
				FoodCalories: []int{5000, 6000},
			},
			{
				FoodCalories: []int{7000, 8000, 9000},
			},
			{
				FoodCalories: []int{10000},
			},
		}

		want := 24000
		got := task1.MaxCalories(exampleElves)
		if got != want {
			t.Fatalf("expected %d, got %d", want, got)
		}
	})
}
