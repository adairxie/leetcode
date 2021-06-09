package main

import "fmt"

func maxProfit(arr []int) int {
	profit := 0
	if len(arr) < 0 {
		return profit
	}

	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) {
			if arr[i] < arr[i+1] {
				profit += arr[i+1] - arr[i]
			}
		}
	}

	return profit
}

func maxProfitV2(prices []int) int {
	var maxProfit, tempProfit int

	for i := 0; i < len(prices); i++ {
		if i+1 < len(prices) {
			tempProfit += prices[i+1] - prices[i]
			if tempProfit <= 0 {
				tempProfit = 0
			}

			if tempProfit > maxProfit {
				maxProfit = tempProfit
			}
		}
	}

	return maxProfit
}

func main() {
	prices := []int{1, 2, 3, 4, 5}
	profit := maxProfitV2(prices)
	fmt.Printf("max_profit: %d\n", profit)
}
