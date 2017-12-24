package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	value, prs := m["k2"]
	fmt.Println(value, prs)
	fmt.Println("Length: ", len(m))

	delete(m, "k2")
	value2, ok := m["k2"]
	fmt.Println(value2, ok)
	fmt.Println("Length: ", len(m))

	someMap := map[int]string{
		0: "Hello",
		1: "Sir",
		2: "Goose",
	}

	if val, exists := someMap[2]; exists {
		fmt.Println("Has value: ", val)
	} else {
		fmt.Println("Does not have value")
	}
	fmt.Println(someMap)
}
