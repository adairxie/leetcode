package main

import "fmt"

func arrayIntersection(a []int, b []int) []int {
	var res []int
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			res = append(res, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 1, 2, 2}
	nums2 := []int{2, 2}

	res := arrayIntersection(nums1, nums2)
	fmt.Printf("%v\n", res)
}
