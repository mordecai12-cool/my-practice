package main

import (
	"fmt"
	"strconv"
	"strings"
)

func marvy(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(up," && i+1 < len(words) {
			result := strings.TrimRight(words[i+1], "),")
			text, err := strconv.Atoi(result)
			if err != nil {
				fmt.Println("Invalid")
			}
			if text > len(words) {
				continue
			}
			for j := 1; j <= text; j++ {
				if i-j >= 0 {
					words[i-j] = strings.ToUpper(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}
	return words
}
func main() {
	fmt.Println(marvy([]string{"this", "is", "so", "fun", "(up,", "2)"}))
}
