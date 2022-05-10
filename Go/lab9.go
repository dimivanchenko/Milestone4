package main

import (
	"fmt"
	//	"log"
	"os"
	"strconv"
)

func main() {
	var folder_path string = os.Args[1]
	var prefix string = os.Args[2]
	counts, _ := strconv.Atoi(os.Args[3])
	//	mode, _ := strconv.Atoi(os.Args[4])
	for i := 0; i < counts; i++ {
		err := os.Mkdir(folder_path+"/"+prefix+strconv.Itoa(i+1), 755)
		if err != nil {
			//		log.Fatal(err)
			fmt.Println("Error. Folder", folder_path+"/"+prefix+strconv.Itoa(i+1), "alredy exist")
		} else {
			fmt.Println("Folder", folder_path+"/"+prefix+strconv.Itoa(i+1), "is createt ")
		}

	}

}
