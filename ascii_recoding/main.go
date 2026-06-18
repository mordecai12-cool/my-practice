package main

import (
	"fmt"
	"os"
	//"strings"
)

func main() {

	
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <text>")
		os.Exit(1)
	}
	input := os.Args[1]

	bannerFile := "standard.txt" // default banner

	if len(os.Args) == 3 {
		validBanner := map[string]bool{
			"standard":   true,
			"shadow":     true,
			"thinkertoy": true,
		}
		if !validBanner[os.Args[2]] {
			fmt.Fprintln(os.Stderr, "Error: unknown banner:", os.Args[2])
			os.Exit(1)
		}
		bannerFile = os.Args[2] + ".txt"
	}

	banner, err := LoadBanner(bannerFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if _, err := ValidateInput(input); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	content := GenerateArt(input, banner)
	fmt.Print(content)
	

}
