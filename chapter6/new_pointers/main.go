package main

import "fmt"

func makeZero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := new(int)
	makeZero(x)
	fmt.Println(*x)
}
