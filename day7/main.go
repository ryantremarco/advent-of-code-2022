package main

import (
	"advent-of-code-2022/m/v2/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	fileTree, err := inputToFileTree(input)
	if err != nil {
		panic(err)
	}

	maxSize := 100000
	totalSize := TotalOfDirectoriesWithinN(fileTree.Node, maxSize)
	fmt.Printf("Total Size within %d: %d\n", maxSize, totalSize)

	filesystemSize := 70000000
	unusedNeeded := 30000000
	deletionCandidateSize := SizeOfDeletionCandidate(fileTree.Node, filesystemSize, unusedNeeded)
	println("Deletion Candidate Size", deletionCandidateSize)
}

// Part 1. Yes, I got bored of putting these in multiple files. Sue me.

func TotalOfDirectoriesWithinN(dir *utils.Node[string, File], maxSize int) int {
	if dir == nil {
		return 0
	}

	totalSize := DirSize(dir)

	totalBeneath := 0
	for _, node := range dir.ChildNodes {
		totalBeneath += TotalOfDirectoriesWithinN(node, maxSize)
	}

	if totalSize > maxSize {
		return totalBeneath
	}

	return totalSize + totalBeneath
}

// Part 2

func SizeOfDeletionCandidate(dir *utils.Node[string, File], filesystemSize, unusedNeeded int) int {
	currentUnused := filesystemSize - DirSize(dir)
	sizeToDelete := unusedNeeded - currentUnused
	candidates := DeletionCandidates(dir, sizeToDelete)
	println("needed", sizeToDelete)
	println("candidates", utils.Map(candidates, func(node *utils.Node[string, File]) int { return DirSize(node) }))
	smallestSize := DirSize(candidates[0])
	for _, candidate := range candidates {
		if size := DirSize(candidate); size < smallestSize {
			smallestSize = size
		}
	}
	return smallestSize
}

func DeletionCandidates(dir *utils.Node[string, File], minSize int) []*utils.Node[string, File] {
	if dir == nil {
		return nil
	}

	childCandidates := []*utils.Node[string, File]{}
	for _, node := range dir.ChildNodes {
		childCandidates = append(childCandidates, DeletionCandidates(node, minSize)...)
	}

	dirSize := DirSize(dir)
	if dirSize >= minSize {
		return append(childCandidates, dir)
	}
	return childCandidates
}

// Common

type File struct {
	Name string
	Size int
}

func DirSize(dir *utils.Node[string, File]) int {
	return utils.Sum(
		utils.Map(
			dir.AllLeaves(),
			func(f File) int {
				return f.Size
			},
		),
	)
}

func inputToFileTree(input string) (utils.Tree[string, File], error) {
	fileTree := utils.NewTree[string, File]("/")
	currentNode := fileTree.Node
	for _, line := range strings.Split(input, "\n")[1:] {
		// Skip ls lines
		if line == "$ ls" {
			continue
		}

		// Swap currentNode on cd
		if strings.HasPrefix(line, "$ cd") {
			dir := strings.Split(line, " cd ")[1]
			switch dir {
			case "..":
				currentNode = currentNode.Parent
			case "/":
				currentNode = fileTree.Node
			default:
				dirNode := currentNode.GetNode(dir)
				if dirNode == nil {
					dirNode = currentNode.AddNode(dir)
				}
				currentNode = dirNode
			}
			continue
		}

		// Create a child node for a dir if we see it
		if strings.HasPrefix(line, "dir ") {
			dir := strings.Split(line, " ")[1]
			currentNode.AddNode(dir)
			continue
		}

		// Here must be a file from an ls. Add it to current node
		fileData := strings.Split(line, " ")
		name := fileData[1]
		size, err := strconv.Atoi(fileData[0])
		if err != nil {
			return utils.Tree[string, File]{}, err
		}
		currentNode.AddLeaf(File{
			Size: size,
			Name: name,
		})
	}

	return fileTree, nil
}
