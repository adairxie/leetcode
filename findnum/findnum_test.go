package findnum

import "testing"

func TestFindNum(t *testing.T) {
	arr := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	if ok := findNum(arr, 5); ok {
		t.Log("find")
	} else {
		t.Log("not found")
	}
}
