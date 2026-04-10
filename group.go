package main

import (
	"fmt"
	"strings"
	"unicode"
)

func capitalize(word string) string {
// type here
	word = strings.TrimSpace(word)
	runes := []rune(strings.ToLower(word))
	runes[2] = unicode.ToUpper(runes[2])
	return string(runes)
}

func main() {
	fmt.Println(capitalize("         learning string as a variable      "))
}

/*package main

import (
	"fmt"
	"strings"
)


func main() {
	s := "MARVY"
	fmt.Println(strings.ToLower(s))

	fmt.Println(strings.ToUpper(`
	1. mango
	2. orange
	3. banana
	4. water melon
	`))

}
*/