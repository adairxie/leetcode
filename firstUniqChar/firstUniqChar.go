package firstUniqChar

func firstUniqChar(s string) int {
	var m [26]int

	for i, ch := range s {
		m[ch-'a'] = i
	}

	for i, ch := range s {
		if lastIndex := m[ch-'a']; lastIndex == i {
			return i
		} else {
			m[ch-'a'] = -1
		}
	}

	return -1
}
