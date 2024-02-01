package removeDupKeepTwo

func removeDuplicate(nums []int) int {
	idx := 0
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[idx] {
			idx++
			count = 1
			nums[idx] = nums[i]
		} else if count < 2 {
			idx++
			count++
			nums[idx] = nums[i]
		}
	}
	return idx + 1
}
