package sword2offer

func FindRepeatNumber(nums []int) int {
	mp := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := mp[n]; ok {
			return n
		}

		mp[n] = struct{}{}
	}
	return -1
}
