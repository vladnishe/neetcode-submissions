func maxArea(heights []int) int {

	l, r := 0, len(heights) - 1
	result := 0
	for l < r {
		w := r - l
	 	h := min(heights[l], heights[r])
		result = max(result, w * h)
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return result
}

func min(h1, h2 int) int {
	if h1 < h2 {
		return h1
	}
	return h2
}

func max(s1, s2 int) int {
	if s1 > s2 {
		return s1
	}
	return s2
}