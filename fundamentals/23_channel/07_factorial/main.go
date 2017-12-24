package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		c := factorial(i)
		for n := range c {
			fmt.Println(i, " ", n)
		}
	}

}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

//func main() {
//	f := factorial(4)
//	fmt.Println("Total: ", f)
//}
//
//func factorial(n int) int {
//	total := 1
//	for i := n; i > 0; i-- {
//		total *= i
//	}
//	return total
//}
