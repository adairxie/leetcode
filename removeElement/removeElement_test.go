package removeElement

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	res := removeElementV2(nums, 2)
	fmt.Printf("%v\n", nums[:res])
	t.Logf("res:%d", res)
}
