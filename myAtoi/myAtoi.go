package myAtoi

import "math"

func myAtoi(str string) int {
    if len(str) == 0 {
		return 0
	}

	positive := 1
	flag := false
	sum := 0
	lastSpaceIndex := -1
	lastNumIndex := -1
	
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if lastSpaceIndex + 1 == i {
				lastSpaceIndex++
				continue
			} else {
				break
			}		
		}

		if str[i] == '-' || str[i] == '+' {
			if flag == false && lastNumIndex == -1 {
				if str[i] == '-' {
					positive = -1
				} else {
					positive = 1
				}
				flag = true
				continue
			}else{
				break
			}
		}

		if str[i] < '0' || str[i] > '9'{
			break
		}

		if lastNumIndex == -1 {
			lastNumIndex = i
		} else if lastNumIndex + 1 == i {
			lastNumIndex++
		} else {
			break
		}

		sum = sum * 10 + int(str[i] - '0')
		if sum * positive > math.MaxInt32 {
			return math.MaxInt32
		}

		if sum * positive < math.MinInt32 {
			return math.MinInt32
		}
	 }

	 return sum * positive
}