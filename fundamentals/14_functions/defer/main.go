package main

import "fmt"

func main() {
	hello()
	world()

	func() {
		fmt.Println("chicken")
	}()

	if getFalse() || getTrue() {
		fmt.Println("in condition")
	}
	if getTrue() || getFalse()  {
		fmt.Println("in condition")
	}
}

func getTrue() bool {
	fmt.Println("In true")
	return true
}

func getFalse() bool {
	fmt.Println("In false")
	return false
}

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("world")
}