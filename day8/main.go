package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	trees, err := inputToTrees(input)
	if err != nil {
		panic(err)
	}

	println("Visibe Trees", countVisibleTrees(trees))
	println("Top Score", topSceneicScore(trees))
}

// part 1
func countVisibleTrees(trees [][]int) int {
	count := len(trees) * len(trees[0])
	for tx := range trees {
		for ty := range trees[tx] {
			tree := trees[tx][ty]
			visibleSides := 4
		dirLoop:
			for _, dir := range []struct {
				x int
				y int
			}{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				for ox, oy := tx+dir.x, ty+dir.y; ox >= 0 && ox < len(trees) && oy >= 0 && oy < len(trees[tx]); ox, oy = ox+dir.x, oy+dir.y {
					otherTree := trees[ox][oy]
					if otherTree >= tree {
						visibleSides--
						continue dirLoop
					}
				}
			}
			if visibleSides == 0 {
				count--
			}
		}
	}
	return count
}

// part 2
func topSceneicScore(trees [][]int) int {
	topScore := 0
	for tx := range trees {
		for ty := range trees[tx] {
			tree := trees[tx][ty]
			vDistances := []int{0, 0, 0, 0}
		dirLoop:
			for di, dir := range []struct {
				x int
				y int
			}{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				for ox, oy := tx+dir.x, ty+dir.y; ox >= 0 && ox < len(trees) && oy >= 0 && oy < len(trees[tx]); ox, oy = ox+dir.x, oy+dir.y {
					otherTree := trees[ox][oy]
					if (otherTree >= tree) || (ox == 0 || ox == len(trees)-1 || oy == 0 || oy == len(trees[tx])-1) {
						vDistances[di] = (tx - ox) + (ty - oy)
						continue dirLoop
					}
				}
			}
			score := 1
			for _, distance := range vDistances {
				if distance < 0 {
					distance = -distance
				}
				score *= distance
			}
			if score > topScore {
				topScore = score
			}
		}
	}
	return topScore
}

func inputToTrees(input string) ([][]int, error) {
	trees := [][]int{}
	for i, line := range strings.Split(input, "\n") {
		treeStrs := strings.Split(line, "")
		trees = append(trees, []int{})
		for _, treeStr := range treeStrs {
			tree, err := strconv.Atoi(treeStr)
			if err != nil {
				return nil, err
			}
			trees[i] = append(trees[i], tree)
		}
	}
	return trees, nil
}