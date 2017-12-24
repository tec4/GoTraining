package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"bufio"
	"log"
	"strings"
)

func main() {
	requiredHeaders := []string{
		"Heading 1",
		"Heading 2",
		"Heading 3",
	}
	processFile("example.csv", requiredHeaders)
}

func processFile(fileName string, requiredHeaders []string)  {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	first := true
	var headers []string
	var headerRowCount int
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if first == true {
			headers, headerRowCount = createHeaders(line, requiredHeaders)
			for i := 0; i < len(headers); i++ {
				fmt.Println(headers[i])
			}
			first = false
			continue
		}

		length := len(line)

		if headerRowCount != length {
			log.Fatalf("Column count mis-match. Headers have %d columns, row has %d columns", headerRowCount, length)
		}

		lineData := make(map[string]string, length)
		for i := 0; i < length; i++ {
			lineData[headers[i]] = strings.TrimSpace(line[i])
		}
		fmt.Println(lineData)
	}
}

// should probably return an error as second argument
func createHeaders(line []string, requiredHeaders []string) ([]string, int) {
	length := len(line)
	var headers []string
	fmt.Println(len(requiredHeaders))
	for i := 0; i < length; i++ {
		headers = append(headers, strings.TrimSpace(line[i]))
	}

	return headers, length
}
