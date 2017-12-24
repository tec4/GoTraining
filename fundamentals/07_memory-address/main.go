package main

import "fmt"

func main() {
	//a := 43
	//fmt.Println("a - ", a)
	//fmt.Println("a's memory adresss is: ", &a)
	//fmt.Printf("a's memory adresss is: %d\n", &a)

	const metersToYards float64 = 1.09361
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, "meters in ", yards, " yards.")

}
