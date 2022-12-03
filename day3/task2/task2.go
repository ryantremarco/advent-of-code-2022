package task2

import (
	"advent-of-code-2022/m/v2/day3/models"
)

func TotalBadgeItemPriorities(rucksacks []models.Rucksack) (int, error) {
	var total int
	groups := make([]models.Group, len(rucksacks)/3)
	for i, rucksack := range rucksacks {
		groups[i/3] = append(groups[i/3], rucksack)
	}
	for _, group := range groups {
		badgeItem, err := group.GetBadgeItem()
		if err != nil {
			return 0, err
		}

		total += badgeItem.Priority()
	}
	return total, nil
}