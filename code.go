package main

import (
	"fmt"
	"strings"
)

func space(s string) string {
	words := strings.Fields(s)
	return strings.Join(words, " ")
}
func main() {
	fmt.Println(space("the  go  main"))
}
