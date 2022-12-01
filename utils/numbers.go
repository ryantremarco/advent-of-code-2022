package utils

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[T Number](nums []T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func Max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

func Min[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return min
}
