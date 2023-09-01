package main

import "fmt"

func makeOddGenerator() func() int {
	i := int(1)
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextOdd := makeOddGenerator()
	for i := 0; i <= 10; i++ {
		fmt.Println(nextOdd())
	}
}
