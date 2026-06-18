package main

import "strings"

func RenderLine(lines string, banner map[rune][]string, colorCode string, letter string) []string {
	var output []string

	for row := 0; row < 8; row++ {
		var result strings.Builder
		for _, ch := range lines {
			ShouldWeColor := colorCode != "" && (letter == "" || strings.ContainsRune(letter, ch))
			if ShouldWeColor {
				result.WriteString(colorCode)
				result.WriteString(banner[ch][row])
				result.WriteString("\033[0m")
			} else {
				result.WriteString(banner[ch][row])
			}
			
		}
		output = append(output, result.String())

	}
	return output
}
