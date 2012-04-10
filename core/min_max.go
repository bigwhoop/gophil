package core

func MinInt(values []int) int {
	var min int
	first := true
	
	for _, value := range values {
		if (first) {
			min = value
			first = false
		} else {
			if value < min {
				min = value
			}
		}
	}
	
	return min
}

func MaxInt(values []int) int {
	var max int
	first := true
	
	for _, value := range values {
		if (first) {
			max = value
			first = false
		} else {
			if value > max {
				max = value
			}
		}
	}
	
	return max
}