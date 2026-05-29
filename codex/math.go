package main

import (
	"fmt"
	"strings"
)

// write a program dat count hw many times each word  appear`s e.g "go go lang exercise exercise"
func countChar(s string) int {
	return len(strings.Fields(s))
}

func main() {
	fmt.Println(countChar("go go"))
	fmt.Println(countChar("lang"))
	fmt.Println(countChar("exercise exercise"))
}
