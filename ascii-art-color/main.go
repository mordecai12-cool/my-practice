// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	color := map[string]string {
// 		"red": "\033[031m",
// 		"green": "\033[032m",
// 		"yellow": "\033[033m",
// 		"blue": "\033[034m",
// 		"purple": "\033[035m",
// 	}
// 	const reset = "\033[0m"

// 	if len(os.Args) != 3 {
// 		fmt.Println("Error: Invalid number of Arguments")
// 		return
// 	}

// 	input := os.Args[1]
// 	text := os.Args[2]

// 	parts := strings.Split(input, "=")
// 	flag := parts[0]
// 	colour := parts[1]

// 	if flag == "--color" {
// 		fmt.Println(color[colour] + text + reset)
// 	} else {
// 		fmt.Println("Invalid flag")
// 		return
// 	}

// }
package main

import (
	"fmt"
	"os"
	"strings"
)

var colourMap = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
}

const reset = "\033[0m"


func main() {
	if len(os.Args) != 3 && len(os.Args) != 4 {
		fmt.Println("Usage:")
		fmt.Println(`go run . --color=<color> "text"`)
		fmt.Println(`go run . --color=<color> <substring> "text"`)
		return
	}

	flag := os.Args[1]
	parts := strings.Split(flag, "=")
	// fmt.Println(parts)

	if len(parts) != 2 || parts[0] != "--color" {
		fmt.Println("Invalid color flag")
		return
	}

	actualColour := parts[1]
	// fmt.Println(actualColour)

	ansiCode, ok := colourMap[actualColour]
	if !ok {
		fmt.Println("Invalid color name")
		return
	}

	if len(os.Args) == 3 {
		text := os.Args[2]
		fmt.Println(ansiCode + text + reset)
		return
	}

	substr := os.Args[2]
	text := os.Args[3]

	var result strings.Builder
	i := 0

	for i < len(text) {
		// fmt.Println(string(text[i]))
		if i+len(substr) <= len(text) &&
			text[i:i+len(substr)] == substr {

				fmt.Println(text[i:i+len(substr)])

			result.WriteString(ansiCode)
			result.WriteString(substr)
			result.WriteString(reset)

			i += len(substr)
		} else {
			result.WriteByte(text[i])
			i++
		}
	}

	fmt.Println(result.String())
	// test()
}