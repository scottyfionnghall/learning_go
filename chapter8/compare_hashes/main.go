package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHash(filename string) (uint32, error) {
	// open the file
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	//create a hasher
	h := crc32.NewIEEE()
	// copy the file into the hasher
	// - copy takes (dst,src) and return (bytesWritten, error)
	_, err2 := io.Copy(h, f)
	// we only need to know of any errors
	if err2 != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

func main() {
	h1, err := getHash("test1.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}
