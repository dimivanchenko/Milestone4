package main

import (
	"fmt"
)

func main() {
	var a, b, c float32
	var min, max float32
	fmt.Print("Enter 3 numbers a b c: ")
	fmt.Scan(&a, &b, &c)
	min = a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	max = a
	if b > max {
		max = b
	}
	if c > min {
		max = c
	}
	fmt.Println("The minimum value among numbers << ", a, b, c, " >> is: ", min)
	fmt.Println("The maximum value among numbers << ", a, b, c, " >> is: ", max)

}
