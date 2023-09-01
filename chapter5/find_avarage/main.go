package main

import "fmt"

func main() {
	numbers := [5]float64{1, 2, 3, 4, 5}
	var total float64 = 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	fmt.Println(total / float64(len(numbers)))
}
