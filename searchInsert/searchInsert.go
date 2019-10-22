package searchInsert

func searchInsert2(nums []int, target int) int {
	low := 0
	high := len(nums)

	for low < high {
		mid := low + (high - low) / 2
		if nums[mid] == target {
			return mid
		} 

		if nums[mid] > target {
			high = mid
		}

		if nums[mid] < target {
			low = mid + 1
		}
	}

	return low
}

func searchInsert(nums []int, target int) int {
	var l = len(nums)
    if l <= 0 {
		return 0
	}

	mid := l / 2

	if nums[mid] == target {
		return mid
	}else if target < nums[mid] {
		return searchInsertCore(nums, target, 0, mid-1)
	} else {
		return searchInsertCore(nums, target, mid+1, l-1)
	}

}

func searchInsertCore(nums []int, target int, start, end int) int {
	if end < start {
		return start
	}

	mid := (start + end) / 2

	if nums[mid] == target {
		return mid
	} else if target < nums[mid] {
		return searchInsertCore(nums, target, 0, mid -1)
	}else{
		return searchInsertCore(nums, target, mid+1, end)
	}
}