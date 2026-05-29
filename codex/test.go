package main

import (
	"fmt"
	"strings"
)

func test(s string) string {
	return strings.ToUpper(s)
}
func main() {
	fmt.Println(test("friend"))
}
