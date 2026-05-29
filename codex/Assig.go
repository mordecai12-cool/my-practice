package main

import (
	"fmt"
	"strconv"
)

func main() {
	var age string
	var option string
	for {
		fmt.Print("Enter age: ")
		fmt.Scanln(&age)
		age, err := strconv.Atoi(age)
		if err != nil {
			fmt.Println("invalid input")
			continue
		} 

		switch {
		case age < 18:
			fmt.Println("child")
		case age >= 18:
			fmt.Println("adult")
		default:
			fmt.Println("invalid input")
			continue
		}
		fmt.Println("choose option(continue, exit): ")
		fmt.Scanln(&option)
		if option == "exit" {
			fmt.Println("good bye!")
			break
		}
	}
}
