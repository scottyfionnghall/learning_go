package main

import "fmt"

func greatestNum(args ...int) int {
	var maxNum int
	for i, v := range args {
		if i == 0 || v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

func main() {
	numbers := []int{5, 2, 4, 3}
	fmt.Println(greatestNum(numbers...))
}
