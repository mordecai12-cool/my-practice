package main

func RenderLine(lines string, banner map[rune][]string) []string {
	result := make([]string, 8)

	for _, ch := range lines {
		for row := 0; row < 8; row++ {
			result[row] += banner[ch][row]
		}
	}
	return result
}
