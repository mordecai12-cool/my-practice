package main

import (
	"fmt"
	"os"
	"strings"
)
func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("invalid!")
		return
	}
	input := strings.Split(os.Args[1], `\n`)

	and := readBanner("standard.txt")
	if len(os.Args) == 3 {
		and = readBanner(os.Args[2] + ".txt")
	}

	for _, ch := range input {
		printBanner(ch, and)
	}
}