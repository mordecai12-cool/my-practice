package main

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	result := make(map[rune][]string)
	for x, key := range base {
		result[x] = key
	}
	for y, col := range priority {
		result[y] = col
	}
	return result
}
