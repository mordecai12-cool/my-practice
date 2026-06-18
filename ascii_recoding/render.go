package main

func RenderLine(lines string, banner map[rune][]string) []string {
	result := make([]string, 8)

	for row := 0; row < 8; row++ {
		for _, ch := range lines {
			result[row] += banner[ch][row]
		}
	}
	return result
}
