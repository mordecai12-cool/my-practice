package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertToDecimal(word []string) string {
	words,err := strconv.ParseInt(word, 16, 64)

	return strings.Join(words, " ")

}
func main() {
	fmt.Println(convertToDecimal([]string{"1E", "(hex)", "files"})) // output: 30 files
}
