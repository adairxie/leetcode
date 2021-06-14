package minPathSum

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	arrs := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	res := minPathSum(arrs)
	t.Logf("min path sum: %d\n", res)

	for i, arr := range arrs {
		for j := 0; j < len(arr); j++ {
			fmt.Printf("%d ", arrs[i][j])
		}
		println()
	}
}
