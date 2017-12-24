package main

import "fmt"

func main() {
	takesVariadicParam(12, 55, 13, 14)
}

func takesVariadicParam(a ...int)  {
	fmt.Printf("%T\n", a)

	var max int
	for _, n := range a {
		fmt.Println(n)
		if n > max {
			max = n
		}
	}

	fmt.Printf("Max value is: %d", max)
}


