// package main

// import (
// 	"strings"
// )

// func GenerateArt(str string, banner map[rune][]string) string {
// 	if str == "" {
// 		return ""
// 	}
// 	parts := SplitInput(str)

// 	onlyEmpty := true

//		for _, part := range parts {
//			if part != "" {
//				onlyEmpty = false
//				break
//			}
//		}
//		var result strings.Builder
//		if onlyEmpty {
//			for i := 0; i < len(parts)-1; i++ {
//				result.WriteString("\n")
//			}
//			return result.String()
//		}
//		for _, part := range parts {
//			if part == "" {
//				result.WriteByte('\n')
//				continue
//			}
//			rows := RenderLine(part, banner)
//			for _, row := range rows {
//				result.WriteString(row)
//				result.WriteString("\n")
//			}
//		}
//		return result.String()
//	}
package main

import (
	"strings"
)

func GenerateArt(text string, banner map[rune][]string, colorCode string, letter string) string {
	if text == "" {
		return ""
	}

	key := SplitInput(text)
	var write strings.Builder
	for i, keys := range key {
		if keys == "" {
			if i < len(key)-1 {
				write.WriteString("\n")
			}
		} else {
			part := RenderLine(keys, banner, colorCode, letter)
			for _, parts := range part {
				write.WriteString(parts)
				write.WriteString("\n")
			}
		}
	}
	return write.String()
}
