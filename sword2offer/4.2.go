package sword2offer

func Search(nums []int, target int) int {
	start := BinSearch(nums, target, 0)
	if start == -1 {
		return 0
	}

	end := BinSearch(nums, target, 1)
	return end - start + 1
}

/*
           /
          /
  A______/
  /      B
 /
/

*/

// last: 0 for A, 1 for B
func BinSearch(nums []int, target int, last int) int {
	if len(nums) < 2 {
		if len(nums) > 0 && nums[0] == target {
			return 0
		}
		return -1
	}

	start, end := 0, len(nums)-1
	for mid := (start + end) / 2; start+1 < end; mid = (start + end) / 2 {
		num := nums[mid]
		switch {
		case num < target:
			start = mid
		case num > target:
			end = mid
		case num == target:
			if last == 0 { // find point A
				if mid-1 < start {
					return mid
				}
				leftOne := nums[mid-1]
				if leftOne < num {
					return mid
				}
				if leftOne == num {
					end = mid
				}
			} else { // find point B
				if mid+1 > end {
					return mid
				}
				rightOne := nums[mid+1]
				if rightOne > num {
					return mid
				}
				if rightOne == num {
					start = mid
				}
			}
		}
	}

	// now start + 1 == end
	if last == 0 {
		if nums[start] == target {
			return start
		}
		if nums[end] == target {
			return end
		}
	} else {
		if nums[end] == target {
			return end
		}
		if nums[start] == target {
			return start
		}
	}

	return -1
}
