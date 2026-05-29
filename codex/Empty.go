package main

import (
	"fmt"
)

func isEmpty(word string) bool {
	return word == ""
}
func main() {
	fmt.Println(isEmpty(""))
	fmt.Println(isEmpty("hello"))
}
