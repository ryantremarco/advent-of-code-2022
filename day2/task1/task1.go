package task1

const winBonus = 6
const drawBonus = 3

func scoreRound(opponent, player int) int {
	score := player
	moveDiff := player - opponent

	if moveDiff == 0 {
		score += drawBonus
	} else if moveDiff == 1 || moveDiff == -2 {
		score += winBonus
	}

	return score
}

var moveMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func TotalScore(rounds [][]string) int {
	total := 0
	for _, round := range rounds {
		opponent := moveMap[round[0]]
		player := moveMap[round[1]]
		total += scoreRound(opponent, player)
	}
	return total
}