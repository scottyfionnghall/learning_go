package main

import (
	"fmt"
	"io"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	urls := []string{
		"http://www.vk.com",
		"http://www.youtube.com",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("The biggest home page:", biggest.URL)
}
