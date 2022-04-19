package sword2offer

import (
	"sort"
)

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	for _, arr := range matrix {
		if len(arr) > 0 && arr[0] <= target && arr[len(arr)-1] >= target {
			idx := sort.SearchInts(arr, target)
			if arr[idx] == target {
				return true
			}
		}
	}
	return false
}
