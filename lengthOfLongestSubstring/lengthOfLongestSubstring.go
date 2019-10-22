package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	var maxLen = 0
	for i := 0; i < len(s); i++ {
		var curLen = 1
		var repeat = make(map[byte]bool)
		for j := i + 1; j < len(s); j++ {
			if _, ok := repeat[s[j]]; ok {
				break
			}
			if s[i] != s[j] {
				repeat[s[j]] = true
				curLen++
			}else{
				break
			}
		}
		
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	
	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	var maxLen = 0
	var flag [256]int
	var beg = 0
	var l = len(s)
	for i := 0; i < l; i++ {
		if flag[s[i]] > 0 && flag[s[i]] > beg {
			beg = flag[s[i]]
		}

		flag[s[i]] = i+1
		maxLen = maxnum(maxLen, i - beg + 1)
	}
	return maxLen
}

func maxnum(a, b int) int {
	if a > b {
		return a
	}

	return b
}