package main

import (
	"fmt"
	"strings"
)

func capLastN(words []string, n int) []string {
	for i := range words {
		if i >= n {
			words[i] = strings.ToUpper(words[i])
		}
	}
	return words
}
func main() {
	fmt.Println(capLastN([]string{"this", "is", "so", "exciting"}, 2))
}
