package stringConvert

import "bytes"

func convert(s string, numRows int) string {
	if numRows <= 0 {
		return ""
	}

	tmp := make([][]byte, numRows)
	var n, incr int = 0, 1

	if numRows == 1 {
		incr = 0
	}
	
	for i := 0; i < len(s); i++ {
		tmp[n] = append(tmp[n], s[i])

		next := n + incr
		if next < 0 || next >= numRows {
			incr = -incr
		}
		n = n + incr
	}

	var buf bytes.Buffer
	for i := range tmp {
		buf.WriteString(string(tmp[i]))
	}

	return buf.String()
}