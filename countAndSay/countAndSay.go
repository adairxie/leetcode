package countAndSay

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	last := countAndSay(n - 1)
	lastBytes := []byte(last)

	var res string
	count := 1
	lastElemt := lastBytes[0]
	for i := 1; i < len(lastBytes); i++ {
		if lastBytes[i] == lastElemt {
			count++
			continue
		} else {
			res = fmt.Sprintf("%s%d%d", res, count, lastElemt-'0')
			count = 1
			lastElemt = lastBytes[i]
		}
	}

	fmt.Println(lastElemt)
	res = fmt.Sprintf("%s%d%d", res, count, lastElemt-'0')

	return res
}
