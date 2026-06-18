package main

import (
	"fmt"
)

func ValidateInput(input string) (rune, error) {
	for _, ch := range input {
		if ch < 32 || ch > 126 {
			return 0, fmt.Errorf("unsupported character %v", ch)
		}
	}
	return 0, nil
}
