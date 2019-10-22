package searchInsert

import "testing"

func TestSearchInsert(t *testing.T) {
	a := []int{}
	res := searchInsert(a, 3)
	t.Logf("res:%d\n", res)
}