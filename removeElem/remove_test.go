package removeElem

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{}
	newLength := removeElement(nums, 2)
	fmt.Printf("new length: %d\n", newLength)
}
