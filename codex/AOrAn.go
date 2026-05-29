package main

import (
	"fmt"
	"strings"
)

func fixAOrAn(word string) string {
	word = strings.ReplaceAll(word, "A", "An")
	word = strings.ReplaceAll(word, "An book", "A book")
	return word
}
func main() {
	fmt.Println(fixAOrAn("there was. A amazing rock. A honest man. A book"))
}
