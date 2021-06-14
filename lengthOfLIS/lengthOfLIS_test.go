package lengthOfLIS

import "testing"

func TestLengthOfLIS(t *testing.T) {
	nums := []int{1, 9, 5, 9, 3}
	l := lengthOfLIS(nums)
	t.Logf("length of LIS: %d\n", l)
}
