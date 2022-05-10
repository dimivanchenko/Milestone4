package main

import (
	"fmt"
)

func minabc(x float32, y float32, z float32) float32 {
	min := x
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}

func maxabc(x float32, y float32, z float32) float32 {
	max := x
	if y > max {
		max = y
	}
	if z > max {
		max = z
	}
	return max
}

func main() {
	var a, b, c float32
	fmt.Print("Enter 3 numbers a b c: ")
	fmt.Scan(&a, &b, &c)
	if minabc(a, b, c) >= -5 && maxabc(a, b, c) <= 5 {
		fmt.Println("OK")
	} else {
		fmt.Println("Wrong")
	}
}
