package main

import (
	"fmt"
	"strings"
)

func fixSingleSpace(word string) string {
	word = strings.Trim(word, "'")
	word = strings.TrimSpace(word)
	return "'" + word + "'"
}
func main() {
	a := "' awesome '"
	b := "' hello word '"
	fmt.Println(fixSingleSpace(a))
	fmt.Println(fixSingleSpace(b))
}
