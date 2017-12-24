package main

import "fmt"

func main() {
	someInt, isEven := getData(22)

	fmt.Println(someInt)
	fmt.Println(isEven)
}

// Get x divided by 2 and determine whether or not the int is even
func getData(x int) (int, bool) {
	return x / 2, x % 2 == 0
}
