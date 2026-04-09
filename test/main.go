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
	first := strings.Fields(words)

	good := strings.Join(first[:1], " ")

	great := strings.Join(first[1:], " ")


	
	fmt.Println(strings.ToUpper(good) + " " + great)

}
