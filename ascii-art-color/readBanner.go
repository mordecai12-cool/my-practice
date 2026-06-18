package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	banner := make(map[rune][]string)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return banner, errors.New("Empty file")
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) < 855 {
		return banner, errors.New("Invalid file")
	}

	for ch := 32; ch <= 126; ch++ {
		start := (ch - 32) * 9
		banner[rune(ch)] = lines[start+1 : start+9]
	}
	return banner, nil
}
