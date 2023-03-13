package main

import (
   "fmt"
   "os"
)

func main() {

   arr := [5]int{1, 2, 3, 4 , 5}

   fmt.Println(os.Args[1:])
   fmt.Println(os.Args[1:2])
   fmt.Println(os.Args[0:len(os.Args)])

   fmt.Println(arr[0:len(arr)])
}