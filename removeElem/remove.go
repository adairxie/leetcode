package removeElem

func removeElement(nums []int, val int) int {
	idx := -1
	for _, num := range nums {
		if num != val {
			idx++
			nums[idx] = num
		}
	}
	return idx + 1
}
