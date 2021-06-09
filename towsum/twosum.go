package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	res := []int{}
	for i, v := range nums {
		dest := target - v
		if index, ok := m[dest]; ok {
			res = append(res, index, i)
			break
		}
		m[v] = i
	}

	return res
}

func main() {
	nums := []int{2, 7, -1, 15}
	target := 6

	res := twoSum(nums, target)
	for _, v := range res {
		fmt.Println(v)
	}
}
