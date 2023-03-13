package main

import (
	"fmt"
)

type address struct {
	street   string
	building int
	pin      int
	city     string
}

type person struct {
	firstName string
	lastName  string
	age       int
	kids      []string
	addr      address
}

func main() {

	sandeepAddr := person{
		firstName: "Sandeep",
		lastName:  "Gholve",
		age:       39,
		kids: []string{
			"Aarohi",
			"Rihan",
		},
		addr: address{
			street:   "Rhinstr",
			building: 161,
			pin:      10315,
			city:     "Berlin",
		},
	}

	fmt.Println(sandeepAddr)

	for _, value := range sandeepAddr.kids {
		fmt.Println(value)
	}

}
