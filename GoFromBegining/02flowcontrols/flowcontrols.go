package main

import (
   "fmt"
   "os"
)

func main() {
	printMessage := len(os.Args) > 1

   if printMessage {
      fmt.Println("Hello ", os.Args[1:])
   }

   booleanAdd()
   conditionOrder()
}

func booleanAdd() {
   year := 2004

   if year % 4 == 0 {
      fmt.Println(year, " is a leap year")
   } else {
      fmt.Println(year, " is not a leap year")
   }

  //fmt.Println(year % 4 == 0 ? "Leap Year" : "Not Leap YEar")
}

func conditionOrder() {
   if (4 + 2) * 10 == 60 {
      fmt.Println("Orders can be defined by parentheses")
   } else {
      fmt.Println("Wrong")
   }
}