package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
