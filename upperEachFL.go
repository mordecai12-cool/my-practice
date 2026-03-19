package main

import (
	"fmt"
	"strings"
)

func upperEachFL(words string) string {
	words = strings.ToLower(words)
	return strings.Title(words)
}
func main() {
	fmt.Println(upperEachFL("hello world main"))
}
