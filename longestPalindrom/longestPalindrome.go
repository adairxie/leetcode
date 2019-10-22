package longestPalindrom


func longestPalindrome(s string) string {	
	ss := preProcess(s)
	var p = make([]int, len(ss))
	var max, id int
	var sLen = len(ss)
	for i := 1; i < sLen - 1; i++ {
		if max > i {
			p[i] = minnum(p[2*id - i], max - i)
		} else {
			p[i] = 1
		}

		for ss[i + p[i]] == ss[i - p[i]] {
			p[i]++
		}

		if max < p[i] {
			max = p[i]
			id = i
		}
	}
	
	var l = max - 1
	var start = (id - l) / 2
	return string(s[start:start+l])
}

func minnum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func preProcess(s string) []byte {
	var l = 2*len(s) + 3
	var ret = make([]byte, l)
	ret[0] = '^'
	for i := 0; i < len(s); i++ {
		ret[2*i+1] = '#'
		ret[2*i+2] = s[i]
	}
	ret[l-2] = '#'
	ret[l-1] = '$'
	
	return ret
}