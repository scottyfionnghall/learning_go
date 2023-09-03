package main

import (
	m "chapter8/math"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Avarage(xs)
	min := m.Min(xs)
	max := m.Max(xs)
	fmt.Println(avg, "\n", min, "\n", max)
}
