package main

import "fmt"

func makeZero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	makeZero(&x)
	fmt.Println(x)
}
