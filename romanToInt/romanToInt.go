package romanToInt

var (
	m = make(map[byte]int, 7)
	sp = make(map[byte]byte, 6)
)

func init() {
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	
	sp['V'] = 'I'
	sp['X'] = 'I'
	sp['L'] = 'X'
	sp['C'] = 'X'
	sp['D'] = 'C'
	sp['M'] = 'C'
}

func romanToInt(s string) int {
	sum := 0
	for i := len(s) -1 ; i >=0 ; i-- {
		sum += m[s[i]]
		if _, ok := sp[s[i]]; ok {
			if ((i - 1 >= 0) && (s[i-1] == sp[s[i]])) {
				sum = sum - m[sp[s[i]]]
				i--
			}
		}
	}

	return sum
}