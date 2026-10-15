func removeDuplicates(nums []int) int {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[p] != nums[i] {
			p += 1
			nums[p] = nums[i]
		}
	}
	return p + 1
}
