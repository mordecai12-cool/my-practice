package main

import (
	"fmt"
	"strings"
)

func correctArticles(sentence string) string {
	word := strings.ReplaceAll(sentence, "a", "an")
	word = strings.ReplaceAll(word, "A", "An")
	return word
}
func cleanQuotes(text []string) string {
	word := strings.Join(text, " ")
	word = strings.ReplaceAll(word, "' ", "'")
	word = strings.ReplaceAll(word, " '", "'")
	return word
}
func main() {
	fmt.Println(correctArticles("It was a unusual event. A honest mistake."))
	fmt.Println(cleanQuotes([]string{"'", "great", "'"}))
	fmt.Println(cleanQuotes([]string{"'", "coding", "is", "fun", "'"}))
}
