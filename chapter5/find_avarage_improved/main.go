package main

import "fmt"

func main() {
	numbers := [5]float64{1, 2, 3, 4, 5}
	var total float64 = 0
	for _, value := range numbers {
		total += value
	}
	fmt.Println(total / float64(len(numbers)))
}
