package main

import "fmt"

func main() {
	switch "Freak" {
	case "Goose1":
		fmt.Println("Goose1")
	case "Goose", "Freak":
		fmt.Println("Goose or Freak")
	default:
		fmt.Println("Default")
	}

	if goose := false; true {
		fmt.Println(goose)
	}

}
