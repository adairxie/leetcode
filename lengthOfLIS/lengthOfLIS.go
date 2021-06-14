package lengthOfLIS

func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	dp := make(map[int]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(dp[i], result)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
