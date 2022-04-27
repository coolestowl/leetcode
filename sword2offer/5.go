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

/*

        /
       /
_ _ _ /       _ _ _
             /
            /
           /



1, 2, 3, 4, 5, 6, 7, 8, 9
6, 7, 8, 9, 1, 2, 3, 4, 5


*/

func MinArray(numbers []int) int {
	min := numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}

func FirstUniqChar(s string) byte {
	counter := make(map[byte]int)
	bts := []byte(s)
	for _, b := range bts {
		counter[b]++
	}

	result := byte(' ')
	for _, b := range bts {
		if counter[b] == 1 {
			result = b
			break
		}
	}
	return result
}
