package main

import (
	"advent-of-code-2022/m/v2/day3/models"
	"advent-of-code-2022/m/v2/day3/task1"
	"advent-of-code-2022/m/v2/day3/task2"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func main() {
	rucksacks, err := inputToRucksacks(input)
	if err != nil {
		panic(err)
	}

	totalMisplacedPriorities := task1.TotalMisplacedPriorities(rucksacks)
	fmt.Println("Total Priorities of Misplaced Items", totalMisplacedPriorities)

	totalBadgeItemPriorities, err := task2.TotalBadgeItemPriorities(rucksacks)
	if err != nil {
		panic(err)
	}
	fmt.Println("Total Badge Item Priorities", totalBadgeItemPriorities)
}

func inputToRucksacks(input string) ([]models.Rucksack, error) {
	lines := strings.Split(input, "\n")
	rucksacks := make([]models.Rucksack, len(lines))
	for i, line := range lines {
		compartmentSize := len(line) / 2
		rucksacks[i] = models.Rucksack{
			CompartmentOne: []models.Item(line[:compartmentSize]),
			CompartmentTwo: []models.Item(line[compartmentSize:]),
		}
	}
	return rucksacks, nil
}
