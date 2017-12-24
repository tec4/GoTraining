package main

import "fmt"

func main() {
	getSumOfMultiplesOf3or5UnderX(1000)
}

func getSumOfMultiplesOf3or5UnderX(n int)  {
	var sum int
	for i := 1; i < n; i++ {
		if i % 5 == 0 || i % 3 == 0 {
			fmt.Printf("Divisable by 3 or 5: %d\n", i)
			sum += i
		}
	}

	fmt.Println(sum)
}



