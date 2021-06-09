package longestCommonPrefix

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	prefix := strs[0]
	for _, k := range strs {
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}

func longestCommonPrefixCore(a, b string) string {
	var i int
	for i = 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			break
		}
	}

	return a[0:i]
}

func longestCommonPrefixV2(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		prefix = longestCommonPrefixCore(strs[i], prefix)
		if prefix == "" {
			break
		}
	}

	return prefix
}
