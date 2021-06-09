package longestCommonPrefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	s := []string{"flower", "flow", "flight"}
	longestCommonPrefix := longestCommonPrefixV2(s)
	t.Logf("the longestCommonPrefix is: %s", longestCommonPrefix)
}
