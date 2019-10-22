package removeElement

import "testing"

func printArray(nums []int, length int) {
	for i := 0; i < length; i++ {
		println(nums[i])
	}
}
func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	res := removeElement(nums, 2)
	printArray(nums, res)
	t.Logf("res:%d", res)
}
