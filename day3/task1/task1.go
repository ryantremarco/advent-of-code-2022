package task1

import (
	"advent-of-code-2022/m/v2/day3/models"

	"golang.org/x/exp/slices"
)

func TotalMisplacedPriorities(rucksacks []models.Rucksack) int {
	var total int
	for _, rucksack := range rucksacks {
		var misplaced models.Item
		for _, char := range rucksack.CompartmentOne {
			if slices.Contains(rucksack.CompartmentTwo, char) {
				misplaced = char
				break
			}
		}
		total += misplaced.Priority()
	}
	return total
}