package sword2offer

func MissingNumber(nums []int) int {
	result := len(nums)
	for i, n := range nums {
		result ^= i ^ n
	}
	return result
}
