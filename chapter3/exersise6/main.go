package main

import "fmt"

func main() {
	var (
		input  float32
		meters float32
	)

	fmt.Print("Enter length in feet: ")
	fmt.Scanf("%f", &input)
	meters = input * 0.3048
	fmt.Println(meters)
}
