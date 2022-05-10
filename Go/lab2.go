package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number string
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	_, err := strconv.Atoi(number)
	if err != nil {
		fmt.Printf("Supplied value << %s >> is not a number\n", number)
	} else {
		fmt.Printf("Supplied value << %s >> is a number\n", number)
	}
}
