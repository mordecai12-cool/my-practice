package main

import "fmt"

func main() {
    // 1. Initialization: i := 1 (Start at 1)
    // 2. Condition: i <= 12 (Stop after 12)
    // 3. Post: i++ (Move to the next number)
    for i := 1; i <= 12; i++ {
        // Use Printf for easy formatting: %d is a placeholder for an integer
        fmt.Printf("3 x %d = %d\n", i, 3*i)
    }
}
