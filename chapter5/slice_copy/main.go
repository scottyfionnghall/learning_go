package main

import "fmt"

func main() {
	numbers := make([]int, 10)
	for i := 0; i <= 9; i++ {
		numbers[i] = i + 1
	}
	fmt.Println(numbers)
}
