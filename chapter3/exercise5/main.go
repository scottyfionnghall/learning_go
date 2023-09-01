// Converts entered temperature in Fharenheit to celsius
package main

import "fmt"

func main() {
	var (
		input   float32
		celsius float32
	)

	fmt.Print("Enter temperature in Faherenheit:")
	fmt.Scanf("%f", &input)
	celsius = (input - 32) * 5 / 9
	fmt.Println(celsius)
}
