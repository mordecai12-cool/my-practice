package main

import (
	"fmt"
	"strings"
)

func letter(test string) string {
	first := strings.ToLower(test[:len(test)-1])
	second := strings.ToUpper(test[len(test)-1:])
	return first + second
}
func main() {
	fmt.Println(letter("world"))
}
