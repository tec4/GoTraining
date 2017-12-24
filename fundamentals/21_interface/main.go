package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Circle struct {
	side float64
}

func (c Circle) area() float64 {
	return math.Pi * c.side * c.side
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
}
