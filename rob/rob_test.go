package rob

import "testing"

func TestRob(t *testing.T) {
	arr := []int{2, 7, 9, 3, 1}
	res := rob(arr)
	t.Logf("%d", res)
}
