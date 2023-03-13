package main

import "fmt"

type Address struct {
   street string
   building int16
   postalCode int16
}



func main() {

   var addresses [5]Address
   addresses[0].street = "Rhinstr"
   addresses[0].building = 161
   addresses[0].postalCode = 10315

   addresses[1] = Address { street: "AlexaStr", postalCode: 12312, building: 121}

   fmt.Println(addresses[0])
   fmt.Println(addresses[1])

   var arrA [5]int
   fmt.Println("emp:", arrA)

   arrA[4] = 100
   fmt.Println("set:", arrA)
   fmt.Println("get:", arrA[4])

   fmt.Println("len:", len(arrA))

   arrB := []int{}

   ints := append(arrB, 10)

   fmt.Println("ints", ints)
   fmt.Printf("%T\n", ints)

   arrB5 := [5]int{1, 2, 3, 4, 5}

   fmt.Printf("%T\n", arrB)
   fmt.Printf("%T\n", arrB5)
   copyOfB := make([]int, len(arrB) + 10)
   fmt.Printf("%T\n", copyOfB)
   copy(copyOfB, arrB)

   fmt.Println("dcl:", arrB)
   fmt.Println("copyOfB:", copyOfB)

   var twoD [2][3]int
   for i := 0; i < 2; i++ {
      for j := 0; j < 3; j++ {
         twoD[i][j] = i + j
      }
   }
   fmt.Println("2d: ", twoD)
}