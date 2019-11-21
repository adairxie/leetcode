package countAndSay

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	res := countAndSay(6)
	t.Logf("output: %s", res)
}
