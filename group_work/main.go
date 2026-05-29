package main

import (
	"fmt"
	"os"
)
func process(data string) string {
	// fmt.Println(strings.ToUpper(string(data))
	//return strings.ToUpper(string(data))
	data = convertCase(data)
	return data
}
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . sample.txt result.txt")
		return
	}
	input := os.Args[1]
	output := os.Args[2]

	//READFILE
	//ARGS
	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	content := string(data)

	result := process(content)
	//result = convertCase(result)
	//result = convertCase(text)

	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		fmt.Println("Erro: unable to write file", err)
		return
	}
}

