package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprint("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	somePeople := []Person{
		{"Bob", 31},
		{"Jimmy", 28},
		{"Zena", 55},
		{"Cowboy", 32},
	}
	fmt.Println(somePeople)
	sort.Sort(ByAge(somePeople))
	fmt.Println(somePeople)

	type people []string
	studyGroup := people{"Aba", "Bambi", "Zaba", "Coono"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)

	group := []string{"Aba", "Bambi", "Zaba", "Coono"}
	sort.Strings(group)
	fmt.Println(group)

	n := []int{7, 4, 8, 2, 9}
	sort.Ints(n)
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

}
