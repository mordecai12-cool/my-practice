package main

import (
	"fmt"
)

func updone(word string) string {
	if len(word) == 0 {

	}
	return string(word[0])
}
func main() {
	fmt.Println(updone("marvellous"))
}
