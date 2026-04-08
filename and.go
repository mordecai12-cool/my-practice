package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hasPrefix(words []string) {
	for i := 0; i < len(words); i++ {

		switch words[i] {
		case "(low,":
			fmt.Println(transform("(low,", words))
		case "(up,":
			fmt.Println(transform("(up,", words))
		case "(cap,":
			fmt.Println(transform("(cap,", words))

		}

	}

}

func main() {
	hasPrefix([]string{"this", "is", "so", "fun", "(cap,", "3)"})
}
func transform(cmd string, words []string) []string {
	count := 1
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], cmd) && i > 0 {
			numStr := strings.TrimRight(words[i+1], ")")
			if n, err := strconv.Atoi(numStr); err == nil {
				count = n
			}
			switch cmd {
			case "(low,":
				for j := i - count; j < i; j++ {
					if j >= 0 {
						words[j] = strings.ToLower(words[j])
					}
				}
			case "(up,":
				for j := i - count; j < i; j++ {
					if j >= 0 {
						words[j] = strings.ToUpper(words[j])
					}
				}
			case "(cap,":
				for j := i - count; j < i; j++ {
					if j >= 0 {
						words[j] = strings.ToUpper(words[j][:1]) + strings.ToLower(words[j][1:])
					}
				}
			}

			words = append(words[:i], words[i+2:]...)
			i--
		}
	}
	return words
}
