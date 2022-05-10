package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b float32
	fmt.Print("Enter 2 numbers a & b (b != 0): ")
	fmt.Scan(&a, &b)
	if b == 0 {
		fmt.Print("Error b = 0 ")
		os.Exit(3)
	}
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)

}
