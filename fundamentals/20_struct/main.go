package main

import "fmt"

type Person struct {
	firstName string
	lastName string
}

type ExtendedPerson struct {
	Person
	firstName string
}

func main() {
	p := ExtendedPerson{
		Person: Person{
			//firstName: "Goose",
			lastName: "Neck",
		},
		firstName: "Tick",
	}
	fmt.Println(p.Person.firstName, p.lastName)
}
