package romanToInt

import "testing"

func TestRomanToInt(t *testing.T) {
	res := romanToInt("MCMXCIV")
	t.Logf("res:%d\n", res)
}