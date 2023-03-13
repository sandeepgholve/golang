package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Failed to read")
	}
	fmt.Println(input)

	// trim the spaces/new line characters
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Grade is ", grade)
	if grade > 60 {
		fmt.Println("Perfect, you passed")
	} else {
		fmt.Println("Ohhh Man, you failed")
	}
}
