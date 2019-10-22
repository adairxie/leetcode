package reverseNum

import "math"

func reverse(x int) int {
	var sum int = 0
	for x != 0 {
		mod := x % 10
		x = x / 10
		if (sum > math.MaxInt32/10 || (sum == math.MaxInt32/10 && mod > 7)) {
			return 0
		}
		if (sum < math.MinInt32/10 || (sum == math.MinInt32/10 && mod < -8)) {
			return 0
		}
		sum = sum * 10 + mod
	}

	return sum
}