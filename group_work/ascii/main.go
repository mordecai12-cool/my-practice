package main

import (
	"fmt"
	"os"
	"strings"
)

// Q1
// func main() {
// 	if len(os.Args) != 2{
// 		fmt.Fprintln(os.Stderr, "Usage: go run main.go <text>")
// 		os.Exit(1)

// 	}
// 	input := os.Args[1]

// 	data, err := os.ReadFile(input)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "Error:", err)
// 		os.Exit(1)
// 	}
// 	content := string(data)
// 	fmt.Println(content)
// }
//Q2
// func main() {
// 	if len(os.Args) != 2{
// 		fmt.Fprintln(os.Stderr, "Usage: go run main.go <text>")
// 		os.Exit(1)
// 	}
// 	input := os.Args[1]
// 	lines := strings.Split(input, ",")

// 	for _, part := range lines {
// 		parts := part
// 		fmt.Println(parts)
// 	}
// 	fmt.Printf("total parts: %d\n", len(lines))
// }
// Q3
// func main() {
// 	if len(os.Args) != 2{
// 		fmt.Fprintln(os.Stderr, "Usage: go run main.go <text>")
// 		os.Exit(1)
// 	}
// 	input := os.Args[1]

// 	for _, i := range input {
// 		fmt.Printf("%c = %d\n",i, i)
// 	}
// }
// Q4
// func main() {
// 	if len(os.Args) != 2{
// 		fmt.Fprintln(os.Stderr, "Usage: go run main.go <text>")
// 		os.Exit(1)
// 	}
// 	input := os.Args[1]

//		for _, i := range input {
//			ch := (i)
//			frontindex := int(ch) - 32
//			fmt.Printf("%c = %d\n", ch, frontindex)
//		}
//	}
//
// Q5
// func main() {
// 	data, err := os.ReadFile(os.Args[1])
// 	if err != nil {
// 		return
// 	}
// 	line := strings.Split(string(data), "\n")
// 	line = line[1:]
	
// 	char := 'H'
// 	ch := int(char-32) * 9
// 	eat := ch + 8
// 	start := line[ch:eat]
// 	//fmt.Println(start)

// 	for _, key := range start {
// 		fmt.Println(key)
// 	}
// }
Q6
func main() {
	
	//line := strings.Split(string(data), "\n")
	
	raw := " _\n| |\n|_|\n \n \n \n \n "
	char := strings.Split(raw, "\n")
	
	//fmt.Println(start)

	for i, key := range char {
		fmt.Println(i, key)
	}
}
