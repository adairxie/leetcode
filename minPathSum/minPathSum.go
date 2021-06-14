package minPathSum

func minPathSum(grid [][]int) int {
	l := len(grid)
	if l < 1 {
		return 0
	}

	dp := grid

	for i := 0; i < l; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
			if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}
			if i != 0 && j != 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[l-1][len(dp[l-1])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
