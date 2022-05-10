package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var address string = os.Args[1]
	var cmd string = os.Args[2]
	var login string = os.Args[3]
	out, err := exec.Command("powershell", "ssh "+login+"@"+address+" "+cmd).Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", out)
	}
}
