package main 
// func dat take a slice and print without space
import (
	"fmt"
	"strings"
)
func cole(word string) string {
	return strings.Join(strings.Fields(word), "")
}
func main() {
	fmt.Println(cole("hello world"))
}