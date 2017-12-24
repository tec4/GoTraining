package main

import "fmt"

func main() {
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)

	letter2 := rune("BA"[0])
	fmt.Println(letter2)
}
