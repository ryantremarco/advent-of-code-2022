package task2

import (
	"advent-of-code-2022/m/v2/day4/models"
)

func OverlappingPairsCount(jobs []models.Job) int {
	var total int
	for i := 0; i < len(jobs); i += 2 {
		job1 := jobs[i]
		job2 := jobs[i+1]

		firstJob := job1
		secondJob := job2

		if firstJob.Start > secondJob.Start {
			firstJob, secondJob = secondJob, firstJob
		}

		if firstJob.End >= secondJob.Start {
			total++
		}
	}
	return total
}