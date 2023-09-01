package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["zero"] = 0
	fmt.Println(x["zero"])
}
