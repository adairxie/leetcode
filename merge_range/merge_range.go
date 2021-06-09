package main

import (
	"fmt"
	"sort"
)

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	res := make([][]int, 0)
	last_front := intervals[0][0]
	last_end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > last_end || intervals[i][1] < last_front {
			res = append(res, []int{last_front, last_end})
			last_front = intervals[i][0]
			last_end = intervals[i][1]
		} else if intervals[i][0] <= last_end {
			last_front = min(last_front, intervals[i][0])
			last_end = max(last_end, intervals[i][1])
		}
	}

	res = append(res, []int{last_front, last_end})
	return res
}

func main() {
	ranges := [][]int{[]int{2, 3}, []int{6, 7}, []int{4, 5}, []int{1, 8}}
	newval := merge(ranges)
	fmt.Printf("%#v", newval)
}
