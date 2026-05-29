package main

//import "fmt"

func StackTwo(top []string, bottom []string) []string {
	result := make([]string, len(top)+len(bottom))
	copy(result, top)
	copy(result[len(top):], bottom)
	return result
}

func StackAll(blocks [][]string) []string {
	code := []string{}
	for _, key := range blocks {
		code = StackTwo(code, key)
	}
	return code
}
// func main() {
// 	fmt.Println(StackTwo([]string{"a","b"}, []string{"c","d"}))
// }
//StackTwo: allocate a new slice of length len(top)+len(bottom), use copy for both halves. StackAll: start with an empty []string{} and call StackTwo(result, block) for each block in the input.