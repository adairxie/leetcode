package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var index int
	var last int = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != last {
			last = nums[i]
			index++
			nums[index] = nums[i]
		}
	}

	return index+1
}

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	len := removeDuplicates(nums)
	for i := 0; i < len; i++ {
		println(nums[i])
	}
}