package main

import (
	"bufio"
	"strings"
	"fmt"
)

func main() {
	const input = "Here is some text."

	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Reading input: ", err)
	}

}
