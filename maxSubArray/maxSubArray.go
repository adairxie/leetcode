package maxSubArray

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	maxSum := nums[0]
	curSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum = curSum + nums[i]
		}

		if curSum > maxSum {
			maxSum = curSum
		}
	}

	return maxSum
}
