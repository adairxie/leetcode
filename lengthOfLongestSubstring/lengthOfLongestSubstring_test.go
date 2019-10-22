package lengthOfLongestSubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	var s = "abcabcbb"
	if lengthOfLongestSubstring(s) != 3 {
		t.Error(`f("abcabcbb") != 3`)
	}

	s = "b"
	if lengthOfLongestSubstring(s) != 1 {
		t.Error(`f("b") != 1`)
	}

	s = "pwwkew"
	if lengthOfLongestSubstring(s) != 3 {
		t.Error(`f("pwwkew") != 3`)
	}
}