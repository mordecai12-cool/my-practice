package main

import (
	"fmt"
	"os"
	//"strings"
)

func main() {
	// 	if len(os.Args) < 2 || len(os.Args) > 3 {
	// 		fmt.Fprintln(os.Stderr, "Usage: go run . <text>")
	// 		os.Exit(1)
	// 	}
	// 	input := os.Args[1]
	// 	//lines := strings.Split(input, "\\n")

	// 	bannerFile := "standard.txt" // default banner

	// 	if len(os.Args) == 3 {
	// 		validBanner := map[string]bool{
	// 			"standard":   true,
	// 			"shadow":     true,
	// 			"thinkertoy": true,
	// 		}
	// 		if !validBanner[os.Args[2]] {
	// 			fmt.Fprintln(os.Stderr, "Error: unknown banner:", os.Args[2])
	// 			os.Exit(1)
	// 		}
	// 		bannerFile = os.Args[2] + ".txt"
	// 	}

	// 	banner, err := loadBanner(bannerFile)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, "Error:", err)
	// 		os.Exit(1)
	// 	}

	// 	if err := validateInput(lines, banner); err != nil {
	// 		fmt.Fprintln(os.Stderr, "Error:", err)
	// 		os.Exit(1)
	// 	}
	// 	RenderLine( banner)
	// }
	// result, err := LoadBanner("standard.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(GenerateArt(os.Args[1], result))
	banner, err := LoadBanner(os.Args[1])
	if err != nil {
		return
	}
	GenerateArt(os.Args[]banner)

}
