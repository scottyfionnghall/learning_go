package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var result int
	for i, v := range x {
		/* we need i == 0, for if we are in 0 postion
		we need to add that value by default */
		if i == 0 || v < result {
			result = v
		}
	}
	fmt.Println(result)
}
