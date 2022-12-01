package utils

import (
	"sort"

	"golang.org/x/exp/constraints"
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
