package main

import (
	"fmt"
	"hash/crc32"
	"os"
)

func main() {
	// create a hasher
	h := crc32.NewIEEE()
	// write our data to it
	file, err := os.ReadFile("test.txt")
	if err != nil {
		return
	}
	h.Write([]byte(file))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)
}
