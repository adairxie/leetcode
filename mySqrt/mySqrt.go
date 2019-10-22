package mySqrt

func mySqrt(x int) int {
	return mySqrtCore(x, 0, x)
}

func mySqrtCore(x, start, end int) int {
	if end * end <= x {
		return end
	}

	mid := (start + end) / 2
	dm := mid * mid
	if dm < x {
		return mySqrtCore(x, mid+1, end) 
	}else if dm > x {
		return mySqrtCore(x, start, mid - 1)
	}

	return mid
}