package main

import (
	"fmt"
	"strings"
)

func ToUpper(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}
func main() {
	fmt.Println(ToUpper([]string{"hello", "world", "fine", "(up)"}))
}
