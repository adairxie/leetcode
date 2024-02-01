package removeDupKeepTwo

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	l := removeDuplicate(nums)
	for i := 0; i < l; i++ {
		fmt.Printf("%d ", nums[i])
	}
}
