package findmediansortedarrays

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums2 := []int{1, 2, 4}
	nums1 := []int{3, 4, 5}
	res := findMedianSortedArrays(nums1, nums2)
	t.Logf("result:%f", res)
}