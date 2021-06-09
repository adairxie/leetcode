package main

import "fmt"

// k may be larger than len(nums)
func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Printf("%v\n", nums)
}
