package main

import (
	"fmt"
	"strings"
)

func updown(word string) string {

	first := strings.ToUpper(string(word[0]))
	middle := strings.ToLower(word[1 : len(word)-1])
	last := strings.ToUpper(word[len(word)-1:])
	return first + middle + last
}
func main() {
	fmt.Println(updown("marvellous"))
}
