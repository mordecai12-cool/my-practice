package main

import (
	"errors"
	"fmt"
)
func ValidateBanner(banner map[rune][]string) error {
	if len(banner)  == 0 {
		return fmt.Errorf("nil")
	}
	if len(banner) != 95  {
		return errors.New("unsupoort char")
	}
	
	for r := rune(32); r <= 126; r++ {
		if len(banner[r]) != 8 {
			return fmt.Errorf("character '%c' has '%v' lines, expected 8", r, len(banner[r]))
		}
	}  
	return nil
}