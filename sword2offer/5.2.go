package sword2offer

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
