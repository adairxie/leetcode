package maxArea

func maxArea(height []int) int {
	maxArea := 0
	i := 0
	j := len(height) - 1
	for i < j {
		lower := i
		high := height[i]
		if high > height[j] {
			high = height[j]
			lower = j
		}
		long := j - i
		area := long * high
		if area > maxArea {
			maxArea = area
		}
		if lower == i {
			i++
		} else {
			j--
		}
	}

	return maxArea
}
