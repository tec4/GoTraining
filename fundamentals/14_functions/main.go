package main

import "fmt"

func main() {
	fmt.Println(getFullNames("hello", "goose"))

	n := average(1,2,3,4,50,22,33)
	fmt.Printf("Average: %.2f \n", n)

	data := []float64{43, 22, 45, 57,23, 44}
	fmt.Printf("TYPE: %T\n", data)
	newAvg := average(data...)
	fmt.Printf("New Average: %.2f \n", newAvg)

	greeting := func() {
		fmt.Println("Hello chicken")
	}

	greeting()

	greet := makeGreeter()
	fmt.Println(greet())

	callCallback([]int{1,2,3,4,5}, func(n int) {
		fmt.Println(n)
	})
}

func getFullNames(firstName, lastName string) (string, string) {
	return firstName + " " + lastName, fmt.Sprint(firstName, " ", lastName)
}

func average(sf ...float64) float64	{
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}

func makeGreeter() func() string {
	return func() string {
		return "Hello greeting which was made"
	}
}

func callCallback(numbers []int, callback func(int))  {
	for _, n := range numbers {
		callback(n)
	}
}