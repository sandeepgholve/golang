package main

import (
	"fmt"
)

// package scope array
var stringArray [5]string

func main() {
	stringArray[0] = "first"
	stringArray[1] = "second"
	stringArray[2] = "third"
	stringArray[3] = "fourth"

	fmt.Println("This is the string array: ", stringArray)

	//Defining and populating an array at function level
	intArray := [5]int{10, 20, 30, 40, 50}

	fmt.Println("This is the int array: ", intArray)

}
