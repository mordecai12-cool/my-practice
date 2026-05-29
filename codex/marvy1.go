package main

import (
	"fmt"
	"strings"
)

func test(s string) string {
	first := strings.ToUpper(s[:1])
	second := s[1:]
	return first + second
}

func good(str string) string{
	middle := strings.ToLower(str[1:len(str)-1])
	first := strings.ToUpper(str[:1])
	second := strings.ToUpper(str[len(str)-1:])
	return first + middle + second
}

func trial(word string) string {
	first := strings.ToUpper(word[:3])
	second := strings.ToLower(word[3:len(word)-1])
	last := strings.ToUpper(word[len(word)-1:])
	return first+second+last
}

func main(){

	fmt.Println(trial("programming")) //PROgramminG
	fmt.Println(good("awesome")) //AwesomE
	fmt.Println(good("congratulation")) //AwesomE
	fmt.Println(test("awesome"))

}

//last character ==> len(var_name) - 1