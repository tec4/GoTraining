package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type personExtended struct {
	person
}

func (p person) fullName() string  {
	return p.firstName + " " + p.lastName
}

func (pe personExtended) fullName() string  {
	return pe.firstName + " ___ " + pe.lastName
}

func main() {
	p1 := person{"Jimmy", "Jones", 20}
	p2 := personExtended{
		person{
			firstName: "Jimmy2",
			lastName: "Jones2",
		},
	}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
