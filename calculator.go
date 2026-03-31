package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter first number: ")
		num1, _ := reader.ReadString('\n')
		num1 = strings.TrimSpace(num1)
		num1Input, err := strconv.Atoi(num1)
		if err != nil {
			fmt.Println("Digit only")
		}
		fmt.Print("select operator(+,-,*,/): ")
		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		fmt.Print("Enter second number: ")
		num2, _ := reader.ReadString('\n')
		num2 = strings.TrimSpace(num2)
		num2Input, err := strconv.Atoi(num2)
		if err != nil {
			fmt.Println("Digit only")
		}

		switch op {
		case "+":
			fmt.Println(num1Input, "+", num2Input, "=", num1Input+num2Input)
		case "-":
			fmt.Println(num1Input, "-", num2Input, "=", num1Input-num2Input)
		case "*":
			fmt.Println(num1Input, "*", num2Input, "=", num1Input*num2Input)
		case "/":
			if num2Input == 0 {
				fmt.Println("indivisible by zero")
				continue
			}
			fmt.Println(num1Input, "/", num2Input, "=", num1Input/num2Input)
		}
		break
	}
}
