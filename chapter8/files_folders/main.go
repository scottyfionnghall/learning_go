package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic("No such file")
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
