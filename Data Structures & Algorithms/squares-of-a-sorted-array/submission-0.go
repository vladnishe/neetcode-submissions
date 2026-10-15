func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	n := len(nums)
	l, r := 0, n - 1
	for i := n - 1; i >= 0; i-- {
		squareLeft := nums[l] * nums[l]
		squareRight := nums[r] * nums[r]
		if squareLeft < squareRight {
			res[i] = squareRight
			r--
		} else {
			res[i] = squareLeft
			l++
		}
	}
	return res
}
