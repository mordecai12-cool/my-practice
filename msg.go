package main

import (
	"fmt"
	"strings"
)

func main() {
	var word, op string
	fmt.Print("Enter word: ")
	fmt.Scanln(&word)

	fmt.Print("Operation (lastChar, capitalize, deleteIndex): ")
	fmt.Scanln(&op)

	switch op {
	case "lastChar":
		fmt.Println(string(word[len(word)-1]))
	case "capitalize":
		fmt.Println(strings.ToUpper(word))
	case "deleteIndex":
		var i int
		fmt.Print("Index: ")
		fmt.Scanln(&i)

		if i < 0 || i >= len(word) {
			fmt.Println("Invalid index")
			return
		}
		fmt.Println(word[:i] + word[i+1:])
	}
}
