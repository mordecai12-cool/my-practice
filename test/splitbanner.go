package main

import "strings"

// Write a function that splits the raw banner text into an array of glyph blocks (each block is 8 lines).

// Ignore trailing empty splits caused by a final separator.
func splitBannerLines(banner string) []string {
	value := strings.TrimSuffix(banner, "\n\n")

	return strings.Split(value, "\n\n")
}
