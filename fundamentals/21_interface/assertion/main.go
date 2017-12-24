package main

import "fmt"

func main() {
	var s interface{} = "Some String"
	str, ok := s.(string)
	if ok {
		fmt.Println(ok, str)
	} else {
		fmt.Println(ok, str)
	}
}
