package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type anotherType struct {
	goosie string
}

type secretAgent struct {
	person
	anotherType
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James"`)
}
func (p anotherType) speak() {
	fmt.Println(`says, "Good morning, James"`)
}

//func (sa secretAgent) speak() {
//	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred"`)
//}

func main() {
	p1 := person{
		"Miss",
		"Moneypenny",
	}

	p2 := secretAgent{
		person{"James", "Bond"},
		anotherType{"chicken"},
		true,
	}

	p1.speak()
	p2.speak()
}