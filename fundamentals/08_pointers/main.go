package main

import "fmt"

func main() {
	fmt.Println(13 % 3)
	//a := 43
	//
	//fmt.Println(a)
	//fmt.Println(&a)
	//
	//var b *int = &a
	//
	//fmt.Println(b)
	//fmt.Println(*b) // dereferencing
	//fmt.Println(b)
	//
	//*b = 42
	//fmt.Println(a)
	x := 5
	fmt.Printf("%p\n", &x)
	zero(&x)
	fmt.Println(x)
}

func zero(z *int) {
	fmt.Printf("%p\n", z)
	*z = 0
}

