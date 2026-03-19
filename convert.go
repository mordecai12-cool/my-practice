package main

import (
	"fmt"
	"strconv"
)

func conv(num string, base int) (int64, error) {
	return strconv.ParseInt(num, base, 64)
}
func main() {
	fmt.Println(conv("1E", 16))
	fmt.Println(conv("1010", 2))
	fmt.Println(conv("77", 8))
}
