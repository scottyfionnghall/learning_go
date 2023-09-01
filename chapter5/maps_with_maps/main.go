package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("What card would you like to see?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	deck := map[string]map[string]string{
		"The Fool": {
			"suit": "Major Arcana",
			"upright_keywords": "beginnings, freedom, innocence, originality, " +
				"adventure, idealism, spontaneity",
		},
	}
	deck["The Fool"]["reversed"] = "reckless, careless, distracted, naive, " +
		"foolish, gullible, stale, dull"
	if card, ok := deck[input]; ok {
		fmt.Println(card["suit"])
	} else {
		fmt.Println("There is no such card")
	}
}
