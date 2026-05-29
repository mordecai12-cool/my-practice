package main

import (
	"fmt"
)

func punc(n string) bool {
	return n == "!" || n == ","
}
func main() {
	fmt.Println(punc("!"))
	fmt.Println(punc(","))
	fmt.Println(punc("x"))
}
