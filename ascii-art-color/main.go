package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	color := map[string]string {
		"red": "\033[031m",
		"green": "\033[032m",
		"yellow": "\033[033m",
		"blue": "\033[034m",
		"purple": "\033[035m",
	}
	const reset = "\033[0m"

	if len(os.Args) != 3 {
		fmt.Println("Error: Invalid number of Arguments")
		return
	}

	input := os.Args[1]
	text := os.Args[2]

	parts := strings.Split(input, "=")
	flag := parts[0]
	colour := parts[1]

	if flag == "--color" {
		fmt.Println(color[colour] + text + reset)
	} else {
		fmt.Println("Invalid flag")
		return
	}

}