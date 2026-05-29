package main 
//func dat return last letter of a word
import "fmt"

func code(word string) string {
	word = word[len(word)-1:]
	return word
}
func main() {
	fmt.Println(code("hello"))
}