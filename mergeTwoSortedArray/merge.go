package mergeTwoSortedArray

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	tail := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[tail] = nums1[i]
			i--
		} else {
			nums1[tail] = nums2[j]
			j--
		}
		tail--
	}

	for j >= 0 && tail >= 0 {
		nums1[tail] = nums2[j]
		j--
		tail--
	}
}
