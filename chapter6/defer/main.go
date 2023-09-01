package main

import "fmt"

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func third() {
	fmt.Println("Third")
}

func main() {
	state := true
	if state {
		defer second()
		third()
	} else {
		defer third()
		second()
	}
	first()
}
