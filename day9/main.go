package main

import (
	"advent-of-code-2022/m/v2/utils"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	moves := inputToMoves(input)

	visitedCount := tailVisitedCount(moves)
	println("Tail Visited Count", visitedCount)

	longVisitedCount := longTailVisitedCount(moves)
	println("Long Tail Visited Count", longVisitedCount)
}

// part 1
func tailVisitedCount(moves []Move) int {
	visited := Points{{0, 0}}
	head := Point{}
	tail := Point{}

	for _, move := range moves {
		for i := 0; i < move.Steps; i++ {
			head.Move(move.Direction)
			xDiff := head.X - tail.X
			yDiff := head.Y - tail.Y

			xDir := Left
			if xDiff >= 0 {
				xDir = Right
			}

			yDir := Down
			if yDiff >= 0 {
				yDir = Up
			}

			if Abs(xDiff) > 1 {
				tail.Move(xDir)
				if yDiff != 0 {
					tail.Move(yDir)
				}
			}
			if Abs(yDiff) > 1 {
				tail.Move(yDir)
				if xDiff != 0 {
					tail.Move(xDir)
				}
			}

			if !visited.Contains(tail) {
				visited = append(visited, tail)
			}
		}
	}

	return len(visited)
}

// part 2
func longTailVisitedCount(moves []Move) int {
	visited := Points{{0, 0}}
	rope := make(Points, 10)

	for _, move := range moves {
		for i := 0; i < move.Steps; i++ {
			rope[0].Move(move.Direction)
			for h := 0; h < len(rope)-1; h++ {
				head, tail := rope[h], rope[h+1]
				xDiff := head.X - tail.X
				yDiff := head.Y - tail.Y

				xDir := Left
				if xDiff >= 0 {
					xDir = Right
				}

				yDir := Down
				if yDiff >= 0 {
					yDir = Up
				}

				if Abs(xDiff) > 1 {
					tail.Move(xDir)
					if yDiff != 0 {
						tail.Move(yDir)
					}
				} else if Abs(yDiff) > 1 {
					tail.Move(yDir)
					if xDiff != 0 {
						tail.Move(xDir)
					}
				}
				rope[h], rope[h+1] = head, tail
			}

			if tail := rope[len(rope)-1]; !visited.Contains(tail) {
				visited = append(visited, tail)
			}
		}
	}

	xMin := utils.Min(utils.Map(visited, func(p Point) int { return p.X }))
	xMax := utils.Max(utils.Map(visited, func(p Point) int { return p.X }))
	yMin := utils.Min(utils.Map(visited, func(p Point) int { return p.Y }))
	yMax := utils.Max(utils.Map(visited, func(p Point) int { return p.Y }))
	for y := yMax; y >= yMin; y-- {
	posLoop:
		for x := xMin; x <= xMax; x++ {
			if x == 0 && y == 0 {
				print("s")
				continue
			}
			for _, point := range visited {
				curr := Point{x, y}
				if point == curr {
					print("#")
					continue posLoop
				}
			}
			print(".")
		}
		println()
	}
	return len(visited)
}

func Abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dir Direction) {
	switch dir {
	case Up:
		p.Y += 1
	case Down:
		p.Y -= 1
	case Left:
		p.X -= 1
	case Right:
		fallthrough
	default:
		p.X += 1
	}
}

type Points []Point

func (p Points) Contains(other Point) bool {
	for _, point := range p {
		if point == other {
			return true
		}
	}
	return false
}

type Direction string

const (
	Up    Direction = "U"
	Down  Direction = "D"
	Left  Direction = "L"
	Right Direction = "R"
)

type Move struct {
	Direction Direction
	Steps     int
}

func inputToMoves(input string) []Move {
	lines := strings.Split(input, "\n")
	moves := make([]Move, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		direction := Direction(parts[0])
		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		moves[i] = Move{
			Direction: direction,
			Steps:     steps,
		}
	}
	return moves
}