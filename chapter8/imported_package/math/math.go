package math

func Avarage(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Max(xs []float64) float64 {
	var maxNum float64
	for i, v := range xs {
		if i == 0 || v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

func Min(xs []float64) float64 {
	var minNum float64
	for i, v := range xs {
		if i == 0 || v < minNum {
			minNum = v
		}
	}
	return minNum
}
