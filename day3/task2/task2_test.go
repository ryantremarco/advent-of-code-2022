package task2_test

import (
	"advent-of-code-2022/m/v2/day3/models"
	"advent-of-code-2022/m/v2/day3/task2"
	"testing"
)

func TestTotalScore(t *testing.T) {
	exampleRucksacks := []models.Rucksack{
		{
			CompartmentOne: []models.Item("vJrwpWtwJgWr"),
			CompartmentTwo: []models.Item("hcsFMMfFFhFp"),
		},
		{
			CompartmentOne: []models.Item("jqHRNqRjqzjGDLGL"),
			CompartmentTwo: []models.Item("rsFMfFZSrLrFZsSL"),
		},
		{
			CompartmentOne: []models.Item("PmmdzqPrV"),
			CompartmentTwo: []models.Item("vPwwTWBwg"),
		},
		{
			CompartmentOne: []models.Item("wMqvLMZHhHMvwLH"),
			CompartmentTwo: []models.Item("jbvcjnnSBnvTQFn"),
		},
		{
			CompartmentOne: []models.Item("ttgJtRGJ"),
			CompartmentTwo: []models.Item("QctTZtZT"),
		},
		{
			CompartmentOne: []models.Item("CrZsJsPPZsGz"),
			CompartmentTwo: []models.Item("wwsLwLmpwMDw"),
		},
	}

	want := 70
	got, err := task2.TotalBadgeItemPriorities(exampleRucksacks)
	if err != nil {
		t.Fatalf("unexpected err %#v", err)
	}

	if want != got {
		t.Fatalf("expected total of %d, got %d", want, got)
	}
}
