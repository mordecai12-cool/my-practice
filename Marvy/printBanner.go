package main

import (
	"fmt"
	"strings"
)

func printBanner(output string, result []string) {
	word := strings.Split(output, "\n")
	for _, marvy := range word {
		if marvy == "" {
			continue
		}

		for i := 0; i < 8; i++ {
			for _, y := range marvy {
				if y >= 32 && y <= 126 {
					good := (int(y) - 32) * 9
					fmt.Print(result[good+1+i])
				}
			}
			fmt.Println()
		}
	}
}
