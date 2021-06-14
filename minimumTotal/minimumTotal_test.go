package minimumTotal

import "testing"

func TestMinimumTotal(t *testing.T) {
	arr := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	result := minimumTotal(arr)
	t.Logf("minimus total path sum: %d\n", result)
}
