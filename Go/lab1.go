package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var question string
	//	var input string
	question = "How are you?"
	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	//	fmt.Println("Данные прочитаны:", answer)
	arr := strings.Split(answer, " ")
	fmt.Println("You are ", arr[2])

}
