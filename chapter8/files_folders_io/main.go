package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	contents, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str := string(contents)
	fmt.Println(str)
}
