package main

import (
	"strings"
)
func convertCase(word string) string {
	words := strings.Fields(word)
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])

			words = append(words[:i], words[i+1:]...)
			i--
		}else if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}else if words[i] == "(cap)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:])
			words = append(words[:i], words[i+1:]...)
			i--
	
		}
	} 
	return strings.Join(words, " ")
}