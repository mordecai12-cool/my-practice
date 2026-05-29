package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1, num2 string
	var op string
	var option string
	for {
		fmt.Print("Enter first number :")
		fmt.Scanln(&num1)
		num1, err := strconv.Atoi(num1)
		if err != nil {
			fmt.Println("Digit only")
			continue
		}
		secondnum:
		fmt.Print("Enter second number :")
		fmt.Scanln(&num2)
		num2, err := strconv.Atoi(num2)
		if err != nil {
			fmt.Println("Digit only")
			goto secondnum
		}

		fmt.Print("operator (+,-,*,/) :")
		fmt.Scanln(&op)
		if num2 == 0 {
			fmt.Println("indivisible by zero")
			continue
		}

		switch op {
		case "+":
			fmt.Println(num1, "+", num2, "=", num1+num2)
		case "-":
			fmt.Println(num1, "-", num2, "=", num1-num2)
		case "*":
			fmt.Println(num1, "*", num2, "=", num1*num2)
		case "/":
			fmt.Println(num1, "/", num2, "=", num1/num2)
		default:
			fmt.Println("invalid operator")
			continue
		}
		fmt.Println("choose option(continue, exit) :")
		fmt.Scanln(&option)
		if option == "exit" {
			fmt.Println("good bye!")
			break
		}
	}
}
