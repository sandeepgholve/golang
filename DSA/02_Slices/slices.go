package main

import (
	"fmt"
	"reflect"
)

// Package scope slices.
var stringSlice []string

// following short notation is not allowed at package level.
// intSlice := []int {10, 20 , 30 , 40}
var intSlice []int = []int{10, 20, 30, 40}

func main() {
	stringSlice = []string{"Sandeep", "Gholve"}

	fmt.Println(intSlice)
	fmt.Println(stringSlice)

	// Define slice using make function
	fractions := make([]float64, 4)
	fractions[0] = 1.1
	fractions[1] = 2.2
	fractions[2] = 3.3
	fractions[3] = 4.4

	fmt.Println(fractions)
	//fractions[4] = 5.5
	fractions = append(fractions, 5.5)
	fmt.Println("fractions type is", reflect.TypeOf(fractions), " and it contains", fractions)
}
