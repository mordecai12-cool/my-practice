package main

import (
	"fmt"
	"strings"
	"unicode"
	"strconv"
)

func capitalize(word string) string {
	words := strings.Fields(word)
	for i, currentWord := range words {
		runes := []rune(strings.ToLower(currentWord))
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}
	words[i] = string(runes)
	}
	return strings.Join(words, " ")
}
func applyCase(word, mode string) string {
	switch mode {
	case "low":
		return strings.ToLower(word)
	case "up":
		return strings.ToUpper(word)
	case "cap":
		return capitalize(word)
	}
	return word

}
func applyMakers(word string) string{
	word = strings.ReplaceAll(word, ", ", ",")
	words := strings.Fields(word)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") {
			inner := strings.Trim(words[i], "()") //"(up,2)" >> up,2
			marvy := strings.Split(inner, ",")//up,2 >> "up", "2"

			mode := marvy[0]//up
			if len(marvy) == 1 {
				if i > 0 {
					words[i-1] = applyCase(words[i-1],mode)
				}
			}else{
				num, err := strconv.Atoi(marvy[1])//2
				if err == nil {
					count := i - num
					if count < 0 {
						count = 0
					}
					
					for j := count; j < i; j++ {
						words[j] = applyCase(words[j], mode)
					}

				}

			}
			words = append(words[:i], words[i+1:]...)
			i--			
		}
	}
	return strings.Join(words, " ")
		
			
}
func main() {
	w := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom,  it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	fmt.Println(applyMakers(w))
}