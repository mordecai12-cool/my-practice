package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	var op string
	for {
		fmt.Println("Enter first number :")
		fmt.Scanln(&num1)

		fmt.Println("Enter second number :")
		fmt.Scanln(&num2)

		fmt.Println("operator (+,-,*,/) :")
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
		}
		break
	}
}
