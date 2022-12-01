package task2

import (
	"advent-of-code-2022/m/v2/day1/models"
	"advent-of-code-2022/m/v2/utils"
)

func SumTopThreeCalories(elves []models.Elf) int {
	calories := make([]int, len(elves))
	for i, elf := range elves {
		calories[i] = utils.Sum(elf.FoodCalories)
	}

	utils.Sort(calories, utils.SortDesc)

	return utils.Sum(calories[0:3])
}