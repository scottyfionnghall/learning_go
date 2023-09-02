package main

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		return
	}
	sha1 := sha1.Sum([]byte(file))
	//os.WriteFile("test.checksum", sha1, 0666)
	fmt.Println(sha1)
}
