package main

import (
	"fmt"
)

// package level map declaration
var nameToPinCodeMap map[string]int

func main() {

	nameToPinCodeMap = map[string]int{
		"Sandeep": 411041,
		"Sagar":   413104,
	}

	for key, value := range nameToPinCodeMap {
		fmt.Println(key, "live in area with pin code", value)
	}

	// function level map short hand declaration
	nameToAgeMap := map[string]int{
		"Sandeep": 39,
		"Milind":  41,
	}

	for key, value := range nameToAgeMap {
		fmt.Println(key, "is", value, "years old.")
	}

	// maps declaration using make function
	countryCodeToCountry := make(map[string]string)
	countryCodeToCountry["INR"] = "India"
	countryCodeToCountry["DEU"] = "Germany"
	countryCodeToCountry["FRA"] = "France"

	for countryCode, country := range countryCodeToCountry {
		fmt.Println(countryCode, " --> ", country)
	}

	delete(countryCodeToCountry, "FRA")
	fmt.Println(countryCodeToCountry)
}
