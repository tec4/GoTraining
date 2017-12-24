package main

import "fmt"

func main() {

	getData := func(x int) (int, bool) {
		return x / 2, x % 2 == 0
	}

	someInt, isEven := getData(22)

	fmt.Println(someInt)
	fmt.Println(isEven)
}
