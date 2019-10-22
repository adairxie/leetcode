package findmediansortedarrays

// l1 = len(nums), l2 = len(nums2)
// target = (l1 + l2) / 2
// todo: find the target th element in nums1 and nums2
// edge conditions:
// 	1. nums1 or nums2 nil
// 实例 1:
// nums1 = [1, 3]
// nums2 = [2]
// 则中位数是 2.0
// 实例 2：
//	nums1 = [1, 2]
//	nums2 = [3, 4]
// 	则中位数是 (2 + 3)/2 = 2.5
// 		
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1 = len(nums1)
	var l2 = len(nums2)
	var total = l1 + l2
	var targetIndex = int(total / 2)
	var merged = make([]int, targetIndex+1)

	var i, j, t int
	for i < l1 && j < l2 {
		if nums1[i] < nums2[j] {
			if t >= targetIndex {
				if total % 2 == 0 {
					return float64(merged[t-1] + nums1[i]) / 2
				}else{
					return float64(nums1[i])
				}
			}
			merged[t] = nums1[i]
			i++
		} else {
			if t >= targetIndex {
				if total % 2 == 0 {
					return float64(merged[t-1] + nums2[j]) / 2
				}else{
					return float64(nums2[j])
				}
			}
			merged[t] = nums2[j]
			j++
		}
		t++
	}

	for ;i < l1; i++ {
		if t >= targetIndex {
			if (total % 2) == 0 {
				return float64(merged[t-1] + nums1[i]) / 2
			}else{
				return float64(nums1[i])
			}
		}
		merged[t] = nums1[i]
		t++
	}

	for ;j < l2; j++ {
		if t >= targetIndex {
			if (total % 2) == 0 {
				return float64(merged[t-1] + nums2[j]) / 2
			}else{
				return float64(nums2[j])
			}
		}
		merged[t] = nums2[j]
		t++
	}

	return float64(0)
}