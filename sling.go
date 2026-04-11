package main

import (
	"fmt"
	// "strings"
)

func dety(name string) bool {

	// var c int16 = -1990-250
	// var d uint16 = 100
	// fmt.Println("int16:", c)
	// fmt.Println("uint16:", d)
	//  var a int = 10
	//  a = 50
	// var b uint = 20
	// fmt.Println("int:", a)
	// fmt.Println("uint:", b)

	// var e float32 = 3.14
	// var f float64 = 6.28
	// fmt.Println("float32:", e)
	// fmt.Println("float64:", f)

	// greeting := "Hello, world"
	// fmt.Println(greeting)

	// 	var status bool = true
	// 	fmt.Println("Status:", status)

	// var age int = 30
	// var pi float64 = 3.14159265359
	// var name string = "John Doe"
	// var isActive bool = true
	// fmt.Printf("Name:, Age:", name, age)
	// fmt.Printf("Value of Pi:", pi)
	// fmt.Printf("Active:", isActive)

	// var name = "Marvy."

	// var salute = `Hi, my name is` + " " + name +" " +` 
	// And I am husting you guys for the weekend`

	// greet := salute

	// fmt.Println(greet)
	if name == "Marvy" {
		return true
	}
	return false

}
func main() {
	fmt.Println(dety("Marvy"))
}