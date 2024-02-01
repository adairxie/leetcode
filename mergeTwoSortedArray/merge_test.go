package mergeTwoSortedArray

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{0}
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Printf("### nums1: %+v\n", nums1)
}
