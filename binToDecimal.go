package main

import (
	"fmt"
	"strconv"
)

func binToDecimal(binStr string) (int64, error) {
	return strconv.ParseInt(binStr, 2, 64)
}
func main() {
	fmt.Println(binToDecimal("10"))
	fmt.Println(binToDecimal("1010"))
}
