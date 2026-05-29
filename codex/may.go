package main 

import (
	"fmt"
	"strconv"
)
func convertNumber(words []string) []string {
	for i := 1; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			number := words[i-1]
			decimal, err := strconv.ParseInt(number, 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(decimal, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}

func convertToDecimal(words []string) []string {
	for i := 1; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			number := words[i-1]
			decimal, err := strconv.ParseInt(number, 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(decimal, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}
func main() {
	fmt.Println(convertNumber([]string{"1E", "(hex)", "files"}))
	fmt.Println(convertToDecimal([]string{"10", "(bin)", "years"}))
}