/* Appends 4,5 to the end of a slice, and returns said slices
 */

package main

import "fmt"

func add_nums(args ...int) []int {
	numbers := args
	numbers = append(numbers, 4, 5)
	return numbers
}

func main() {

	numbers := []int{1, 2, 3}
	fmt.Println(add_nums(numbers...))
}
