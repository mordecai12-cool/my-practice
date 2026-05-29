package main

import (
	"fmt"
	"strings"
)

func capFirstLetter(words string) string {
	return strings.ToUpper(words[:1]) + strings.ToLower(words[1:])
}
func main() {
	fmt.Println(capFirstLetter("hELLO"))
	fmt.Println(capFirstLetter("WORLD"))
	fmt.Println(capFirstLetter("aWeSoMe"))
}
