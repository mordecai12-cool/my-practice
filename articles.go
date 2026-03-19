package main

import (
	"fmt"
	"strings"
)

func fixArticles(s []string) string {
	word := strings.Join(s, " ")
	word = strings.ReplaceAll(word, " !", "!")
	word = strings.ReplaceAll(word, " ,", ",")
	return word
}
func main() {
	fmt.Println(fixArticles([]string{"hello", ",", "world", "!"}))
}
