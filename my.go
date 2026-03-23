package main

import (
	"fmt"
	"strings"
)

func my(word string) string {
	if len(word) == 0 {

	}
	return strings.ToUpper(word[len(word)-1:])
}
func main() {
	fmt.Println(my("friend"))
}
