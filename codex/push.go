package main 

import (
	"fmt"
	"strings"
	//"strconv"
)
func ToupperLastWord(words []string) []string {
	for i := 1; i < len(words); i++ {
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			
			words = append(words[:i], words[i+1:]...)
			i--
		} 
	} 
	return words
}
func main() {
	fmt.Println(ToupperLastWord([]string{"hello", "world", "(up)"}))
}