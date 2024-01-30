package findnum

func findNum(array [][]int, target int) bool {
	length := len(array)
	if length == 0 {
		return false
	}
	width := len(array[0])
	if width == 0 {
		return false
	}

	i := 0
	j := length - 1
	for i < width && j >= 0 {
		if array[i][j] == target {
			return true
		} else if array[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
