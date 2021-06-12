package maxSubArray

import "testing"

func TestMaxSubArray(t *testing.T) {
	nums := []int{-1, -2, 2, -1, 3}
	max := maxSubArray(nums)
	t.Logf("max sub sum: %d\n", max)
}
