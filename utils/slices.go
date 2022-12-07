package utils

import (
	"sort"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type SortDirection string

const (
	SortAsc  = "ASC"
	SortDesc = "DESC"
)

func Sort[T constraints.Ordered](arr []T, dir SortDirection) {
	sort.Slice(arr, func(i, j int) bool {
		switch dir {
		case SortDesc:
			return arr[i] > arr[j]
		case SortAsc:
			fallthrough
		default:
			return arr[i] < arr[j]
		}
	})
}

func Remove[T comparable](arr []T, toRemove T) []T {
	copy := append([]T{}, arr...)

	for slices.Contains(copy, toRemove) {
		for i, item := range copy {
			if item == toRemove {
				copy = append(copy[:i], copy[i+1:]...)
				break
			}
		}
	}

	return copy
}

func Map[T any, V any](arr []T, mapper func(T) V) []V {
	out := make([]V, len(arr))
	for i, item := range arr {
		out[i] = mapper(item)
	}
	return out
}
