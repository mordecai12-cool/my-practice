package main

import (
	"fmt"
	"strings"
)

func main() {
	var w, op string
	fmt.Scanln(&w, &op)

	fmt.Print("Operation (lastChar, capitalize, deleteIndex): ")
	fmt.Scanln(&op)

	if op == "lastChar" {
		fmt.Println(string(w[len(w)-1]))
	} else if op == "capitalize" {
		fmt.Println(strings.ToUpper(w))
	} else if op == "deleteIndex" {
		var i int
		fmt.Scanln(&i)
		if i >= 0 && i < len(w) {
			fmt.Println(w[:i] + w[i+1:])
		}
	}
}
