package task1

import (
	"advent-of-code-2022/m/v2/day1/models"
	"advent-of-code-2022/m/v2/utils"
)

func MaxCalories(elves []models.Elf) int {
	calories := make([]int, len(elves))
	for i, elf := range elves {
		calories[i] = utils.Sum(elf.FoodCalories)
	}
	return utils.Max(calories)
}