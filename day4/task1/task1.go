package task1

import (
	"advent-of-code-2022/m/v2/day4/models"
)

func ContainedPairsCount(jobs []models.Job) int {
	var total int
	for i := 0; i < len(jobs); i += 2 {
		job1 := jobs[i]
		job2 := jobs[i+1]

		firstStart := job1.Start
		if job2.Start <= firstStart {
			firstStart = job2.Start
		}

		lastEnd := job1.End
		if job2.End >= lastEnd {
			lastEnd = job2.End
		}

		combinedJob := models.Job{Start: firstStart, End: lastEnd}

		if combinedJob == job1 || combinedJob == job2 {
			total++
		}
	}
	return total
}