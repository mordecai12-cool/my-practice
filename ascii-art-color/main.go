package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	color := flag.String("color", "", "red, green, yellow, ...")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 || len(args) > 3 {
		fmt.Fprintln(os.Stderr, "usage: go run . --color=<color> [letters] text [banner]")
		os.Exit(1)
	}

	var letter, input string
	bannerFile := "standard.txt"

	switch len(args) {
	case 1:
		input = args[0]
	case 2:
		letter = args[0]
		input = args[1]
	case 3:
		letter = args[0]
		input = args[1]
		bannerFile = args[2] + ".txt"
	}
	colorCode := ""

	if *color != "" {
		code, ok := ColorMap[*color]
		if !ok {
			fmt.Fprintln(os.Stderr, "not a color")
		}
		colorCode = code
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
	content := GenerateArt(input, banner, colorCode, letter)
	fmt.Print(content)

}
