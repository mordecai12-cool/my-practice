package main 
// func dat conv "10" to num 10
import (
	"fmt"
	"strconv"
)
func convertToTen(word string) (int, error) {
	return strconv.Atoi(word)
}

func main() {
	fmt.Println(convertToTen("10"))
}