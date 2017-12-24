package main

import "fmt"

func main() {
	x := 0

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println()

	increment2 := wrapper()
	fmt.Println(increment2())
	fmt.Println(increment2())
}

func wrapper() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}
