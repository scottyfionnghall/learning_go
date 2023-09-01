package main

import "fmt"

func avarage(numbers []float64) (float64, bool) {
	total := 0.0
	for _, value := range numbers {
		total += value
	}

	return total / float64(len(numbers)), true
}

func main() {
	numbers := []float64{98, 93, 77, 82, 83}
	result, ok := avarage(numbers)
	if ok {
		fmt.Println(result)
	}
}
