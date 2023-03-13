package main

import (
   "fmt"
   "os"
)

func main() {
   fmt.Println("Hello Go!")

   var name string = "Sandeep Gholve"
   city := "Berlin"

   var (
      street = "Rhinstr"
      building = 161
      postalCode = 10315
   )

   var echo = fmt.Printf

   fmt.Println("I am ", name, " and I live in ", city, " and My address is ", street, building, postalCode)
   _, err := echo("I am %s and I live in %s. And My detailed address is: %s %d, %d\n", name, city, street, building, postalCode, postalCode)
   fmt.Println("\nBefore Error Check = ")
   if err != nil {
      fmt.Println("ERROR1 = ")
      fmt.Println(err)
      fmt.Println("ERROR2 = ")
   } else {
      fmt.Println("In Else")
   }
   fmt.Println("After Error Check.")
   fmt.Println("\nCommand line argument passed is --> ", os.Args)
}


/*
Go Run.json config

{
   "name": "go build sandeep.com/hello",
   "type": "go",
   "goExecPath": "/usr/local/go/bin/go",
   "buildParams": ["sandeep.com/hello"],
}
*/

