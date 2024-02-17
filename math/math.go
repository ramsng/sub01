package math

func Min(xa []float64) float64 {
	var min float64 = 0
	for ind, val := range xa {
		if ind == 0 {
			min = val
		} else {
			switch {
			case min <= val:
				break
			case min > val:
				min = val
			}
		}
	}
	return min
}

func Max(xa []float64) float64 {
	var max float64 = 0
	for ind, val := range xa {
		if ind == 0 {
			max = val
		} else {
			switch {
			case max >= val:
				break
			case max < val:
				max = val
			}
		}
	}
	return max
}
