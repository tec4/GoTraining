package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v - type: %T \n", a, a)
	fmt.Printf("%v - type: %T  \n", b, b)
	fmt.Printf("%v - type: %T  \n", c, c)
	fmt.Printf("%v - type: %T  \n", d, d)


	fmt.Println()

	// Zero values
	var e int
	var f string
	var g float64
	var h bool

	fmt.Printf("%v \n", e)
	fmt.Printf("'%v' \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)
	fmt.Println()
}
