package main

import "fmt"

func plusone(digits []int) []int {
	if len(digits) < 1 {
		return []int{1}
	}

	var carry int
	if digits[len(digits)-1]+1 >= 10 {
		digits[len(digits)-1] = 0
		carry = 1
	} else {
		digits[len(digits)-1]++
		return digits
	}

	for i := len(digits) - 2; i >= 0; i-- {
		if digits[i]+carry < 10 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
			carry = 1
		}
	}

	var res = []int{carry}
	res = append(res, digits...)
	return res
}

func main() {
	arr := []int{1, 3, 9, 8, 9}
	narr := plusone(arr)
	fmt.Printf("%v\n", narr)
}
