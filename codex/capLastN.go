package main

import (
	"fmt"
	"strings"
)

// func capLastN(words []string, n int) []string {
// 	for i := range words {
// 		if i >= n {
// 			words[i] = strings.ToUpper(words[i])
// 		}
// 	}
// 	return words
// }

func uppercaseLastWords(words []string, count int) []string {
  // write your code here
  startIndex := len(words) - count
  for i := range words {
    if i >= startIndex {
       words[i] = strings.ToUpper(words[i])
    }
  }
  return words
}
func main() {
	fmt.Println(uppercaseLastWords([]string{"this", "is", "so", "exciting"}, 1))
}
