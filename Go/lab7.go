package main

import (
	"fmt"
	"os"
)

func basicOperations(x float32, y float32) {
	fmt.Println("a + b = ", x+y)
	fmt.Println("a - b = ", x-y)
	fmt.Println("a * b = ", x*y)
	fmt.Println("a / b = ", x/y)
}

func main() {
	var a, b float32
	fmt.Print("Enter numbers a & b (b != 0): ")
	fmt.Scan(&a, &b)
	if b == 0 {
		fmt.Print("Error b = 0 ")
		os.Exit(3)
	}
	basicOperations(a, b)
}
