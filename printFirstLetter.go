package main

import (
	"fmt"
)

func firstLetter(word string) string {
	if len(word) == 0 {
	}
	return string(word[0])
}
func main() {
	fmt.Println(firstLetter("hello"))
}
