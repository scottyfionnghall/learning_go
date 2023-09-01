package main

import "fmt"

func makeEvenGenerator() func() int {
	i := int(0)
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextEven := makeEvenGenerator()
	for i := 0; i <= 10; i++ {
		fmt.Println(nextEven())
	}
}
