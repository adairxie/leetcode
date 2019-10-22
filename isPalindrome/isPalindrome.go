package isPalindrome

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var t = 0
	for x > t {
		t = t*10 + x%10
		x = x / 10
	}

	return x == t || x == t/10
}
