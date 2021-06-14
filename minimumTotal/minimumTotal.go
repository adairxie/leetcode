package minimumTotal

func minimumTotal(triangle [][]int) int {
	if len(triangle) < 1 {
		return 0
	}

	result := 1<<31 - 1
	dp := map[int]map[int]int{}
	dp[0] = map[int]int{}
	dp[0][0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		dp[i] = map[int]int{}
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		var j = 1
		for j = 1; j < len(triangle[i])-1; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}

		dp[i][j] = dp[i-1][j-1] + triangle[i][j]
	}

	last := len(triangle) - 1
	for j := 0; j < len(dp[last]); j++ {
		result = min(dp[last][j], result)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
