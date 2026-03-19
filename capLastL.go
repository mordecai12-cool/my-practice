package main

import (
	"fmt"
	"strings"
)

func capLastL(s string) string {
	word := strings.Fields(strings.ToLower(s))
	for i := range word {
		word[i] = word[i][:len(word[i])-1] + strings.ToUpper(word[i][len(word[i])-1:])
	}
	return strings.Join(word, " ")
}
func main() {
	fmt.Println(capLastL("Hello World Main"))
}
