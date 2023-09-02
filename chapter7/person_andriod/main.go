package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p Person) talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	andi := Android{
		Model: "dick",
		Person: Person{
			Name: "Andi",
		},
	}
	tom := Person{Name: "Tom"}
	andi.talk()
	tom.talk()
}
