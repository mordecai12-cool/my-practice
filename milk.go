package main

import (
	"fmt"
	"strings"
)

/*func space(s string) string {
	word := strings.ToUpper(s)
	return word
}
*/

func main() {
	var word string
	fmt.Println("Enter word")
	fmt.Scanln(&word)

	fmt.Println(strings.ToLower(string(word[:len(word)-1])) + strings.ToUpper(word[len(word)-1:]))
}
