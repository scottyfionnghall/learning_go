package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		defer func() {
			str := recover()
			fmt.Println(str)
		}()
		panic("Could not create file")
	}
	defer file.Close()

	file.WriteString("test 2")
}
