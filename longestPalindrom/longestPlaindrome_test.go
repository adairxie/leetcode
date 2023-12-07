package longestPalindrom

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	res := longestPalindrome("abbc")
	t.Logf("output:%s\n", res)
}
