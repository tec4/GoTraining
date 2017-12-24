package main

import "fmt"

func main() {
	//mySlice := []string{"a", "b", "c", "d", "e", "f", "g"}
	//fmt.Println(mySlice)
	//fmt.Println(mySlice[2:4])
	//fmt.Println(mySlice[2])
	//fmt.Println("myString"[2])
	//fmt.Println(string("myString"[2]))

	//mySlice2 := make([]int, 0, 5)
	//fmt.Println("--------------")
	//fmt.Println(mySlice2)
	//fmt.Println(len(mySlice2))
	//fmt.Println(cap(mySlice2))
	//fmt.Println("--------------")
	//
	//for i := 0; i < 80; i++ {
	//	mySlice2 = append(mySlice2, i)
	//	fmt.Println("Len: ", len(mySlice2), "Capacity: ", cap(mySlice2), "Value: ", mySlice2[i])
	//}

	//greeting := make([]string, 2, 5)
	//greeting[0] = "Turkey"
	//greeting = append(greeting, "Turkey", "Chicken", "Freak")

	//for k, v := range greeting {
	//	fmt.Println(k, v)
	//}
	//
	//fmt.Println(len(greeting), cap(greeting))


	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++
	fmt.Println(mySlice[0])

}
