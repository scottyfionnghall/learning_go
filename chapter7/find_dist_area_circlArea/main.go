package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	width, height float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

func (rec Rectangle) area() float64 {
	area := rec.width * rec.height
	return area
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (rec Rectangle) perimeter() float64 {
	perimeter := 2 * (rec.width + rec.height)
	return perimeter
}

func (c Circle) perimeter() float64 {
	perimeter := 2 * math.Pi * c.r
	return perimeter
}

func (rec Rectangle) diagonal() float64 {
	diagonal := math.Sqrt(math.Pow(rec.width, 2) + math.Pow(rec.height, 2))
	return diagonal
}

func findAllArea(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Println(s.area())
	}
}

func findAllPerimeter(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Println(s.perimeter())
	}
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println("Circle with a radius of", c.r, "has an area of", c.area())
	rec := Rectangle{width: 5, height: 3}
	fmt.Println("Rectangle with a width of", rec.width, "and height of",
		rec.height, "has an area of", rec.area())
	findAllArea(c, rec)
	findAllPerimeter(c, rec)
}
