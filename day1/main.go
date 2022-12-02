package main

import (
	"advent-of-code-2022/m/v2/day1/models"
	"advent-of-code-2022/m/v2/day1/task1"
	"advent-of-code-2022/m/v2/day1/task2"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	elves, err := inputToElves(input)
	if err != nil {
		panic(err)
	}

	maxCalories := task1.MaxCalories(elves)
	sumTopThreeCalories := task2.SumTopThreeCalories(elves)

	fmt.Println("Max Calories", maxCalories)
	fmt.Println("Sum Top Three Calories", sumTopThreeCalories)
}

func inputToElves(input string) ([]models.Elf, error) {
	lines := strings.Split(input, "\n")
	elves := []models.Elf{{}}
	currentElf := 0
	for _, line := range lines {
		if line == "" {
			elves = append(elves, models.Elf{})
			currentElf++
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("failed to convert %s to int", line))
		}
		elves[currentElf].FoodCalories = append(elves[currentElf].FoodCalories, calories)
	}
	return elves, nil
}