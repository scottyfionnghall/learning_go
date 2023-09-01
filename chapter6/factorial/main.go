package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	fmt.Println(x)
	return x * factorial(x-1)
}

func main() {
	var input uint
	input = 10
	fmt.Println("Factorial of ", input, " is ", factorial(input))
}
