package main

import (
	"fmt"
	"strings"
)

func main() {
	var word, op string
	fmt.Println("enter word:")
	fmt.Scanln(&word)
	for {
		fmt.Println("operation (lastChar, capitalize, deleteIndex): ")
		fmt.Scanln(&op)

		switch op {
		case "lastChar":
			fmt.Println(string(word[len(word)-1]))
		case "capitalize":
			fmt.Println(strings.ToUpper(word))
		case "deleteIndex":
			var i int
			fmt.Println("Enter index:")
			fmt.Scanln(&i)
			if i < 0 || i >= len(word) {
				fmt.Println("invalid index")
				main()
				return
			}
			fmt.Println(word[:i] + word[i+1:])
		}
	}
}
