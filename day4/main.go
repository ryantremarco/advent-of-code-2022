package main

import (
	"advent-of-code-2022/m/v2/day4/models"
	"advent-of-code-2022/m/v2/day4/task1"
	"advent-of-code-2022/m/v2/day4/task2"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	jobs, err := inputToJobs(input)
	if err != nil {
		panic(err)
	}

	containedPairs := task1.ContainedPairsCount(jobs)
	println("Number Contained Pairs", containedPairs)

	overlappingPairs := task2.OverlappingPairsCount(jobs)
	println("Number of Overlapping Pairs", overlappingPairs)
}

func inputToJobs(input string) ([]models.Job, error) {
	lines := strings.Split(strings.ReplaceAll(input, ",", "\n"), "\n")
	jobs := make([]models.Job, len(lines))
	for i, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Split(line, "-")

		start, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}

		end, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}

		jobs[i] = models.Job{
			Start: start,
			End:   end,
		}
	}
	return jobs, nil
}
