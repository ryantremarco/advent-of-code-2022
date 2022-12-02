package task2

const winBonus = 6
const drawBonus = 3

func scoreGuidedRound(opponent int, result string) int {
	var score int
	var moveOffset int

	switch result {
	case "Z":
		// win
		score += winBonus
		moveOffset = 1
	case "Y":
		// draw
		score += drawBonus
		moveOffset = 0
	case "X":
		fallthrough
	default:
		// lose
		moveOffset = -1
	}

	move := (opponent + moveOffset)
	for move > 3 {
		move = move - 3
	}
	for move < 1 {
		move = move + 3
	}
	score += move

	return score
}

var moveMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func GuidedTotalScore(rounds [][]string) int {
	total := 0
	for _, round := range rounds {
		opponent := moveMap[round[0]]
		guide := round[1]
		total += scoreGuidedRound(opponent, guide)
	}
	return total
}