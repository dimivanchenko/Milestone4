package main

import (
	"fmt"
	"strings" // Needed to use Split
)

func main() {
	//	str := "teststring"
	var str string
	fmt.Print("Input string: ")
	fmt.Scan(&str)
	split := strings.Split(str, "")
	fmt.Println(split)
	//	fmt.Println("The length of the slice is:", len(split))
	//	for i := 0; i < len(split); i++ {}
	for idx, char := range split {
		fmt.Printf("Char %d is: %s\n", idx, char)
	}

}
