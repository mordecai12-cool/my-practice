package main 
// func dat check if 2words re desame e.g ("hi", "hi"), ("hi", "hello")
import "fmt"

func code(word1, word2 string) bool {
	return word1 == word2
}

func main() {
	fmt.Println(code("hi", "hi"))
	fmt.Println(code("hi", "hello"))
}