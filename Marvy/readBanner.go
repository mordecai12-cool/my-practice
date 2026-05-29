package main

import (
	"fmt"
	"os"
	"strings"
)

func readBanner(input string) []string {
	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error :")
		return nil
	}
	return strings.Split(string(data), "\n")
}
