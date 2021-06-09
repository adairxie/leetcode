package removeElement

func removeElement(nums []int, val int) int {
	var start int = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		} else {
			start++
			nums[start] = nums[i]
		}
	}

	return start + 1
}

func removeElementV2(nums []int, val int) int {
	toDeleteIndx := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			toDeleteIndx++
			nums[toDeleteIndx] = nums[i]
		}
	}

	return toDeleteIndx + 1
}
