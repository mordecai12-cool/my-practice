package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	learn := (os.Args[1])
	result, err := os.ReadFile(learn)
	if err != nil {
		fmt.Println("invalid")
	}
	words := string(result)

	fmt.Println(strings.ToUpper(words))
	fmt.Println(strings.Title(words))
	fmt.Println(words[0:11])
	fmt.Println(words[55:])

}
