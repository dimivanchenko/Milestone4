package main

import (
	"fmt"
)

func main() {
	var a [100]int
	var count, element int
	var max, min int
	fmt.Print("enter the number of array elements: ")
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Print("enter array ", i, " element: ")
		fmt.Scan(&element)
		a[i] = element
	}
	for i := 0; i < count; i++ {
		fmt.Println("element ", i, " is ", a[i])

	}
	min = a[0]
	max = a[0]
	for i := 0; i < count; i++ {
		if a[i] < min {
			min = a[i]
		}
		if a[i] > max {
			max = a[i]
		}

	}

	fmt.Println("The minimum value of an array element is: ", min)
	fmt.Println("The maximum value of an array element is: ", max)
}
