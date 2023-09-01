package main

import "fmt"

func half(number int) (float64, bool) {
	number = number / 2
	if number%2 == 0 {
		return float64(number), true
	} else {
		return float64(number), false
	}
}

func main() {
	number := 7
	fmt.Println("Is", number, "even?")
	result, ok := half(number)
	if ok {
		fmt.Println("Yes,", number, "is an even number.")
	} else {
		fmt.Println("No,", number, "is an odd number.")
	}
	fmt.Println("Also,", number, "in half is", result)
}
