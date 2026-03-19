package main

import (
	"fmt"
	"strings"
)

func vowel(word string) string {
	if strings.ContainsAny(string(word[0]), "aeiouhAEIOUH") {
		return "an"
	}
	return "a"
}
func main() {
	fmt.Println(vowel("apple"))
	fmt.Println(vowel("book"))
	fmt.Println(vowel("egg"))
	fmt.Println(vowel("horse"))
}
