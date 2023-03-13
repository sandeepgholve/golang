package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	var price int = 100
	fmt.Println("Price is ", price, " dollars.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is ", tax, " dollars.")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is ", total, " dollars.")

	var availableFunds int = 120
	fmt.Println("Available funds are ", availableFunds, " dollars.")
	fmt.Println(" Within budget? ", total <= float64(availableFunds))

	//testBasics()
	var originalCount int = 10
	var eatenCount int = 4

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")

}

func testBasics() {
	trueBoolean := true
	falseBoolean := false
	str := strings.Title("head first go")
	f := math.Floor(2.70)

	var quantity int
	var length, width float64
	var customerName string
	var flag bool

	quantity1 := 10
	//quantity1 := 20

	fmt.Println(quantity1)

	fmt.Println("Hello, Go", "Sandeep Gholve", "Berlin")
	fmt.Println(str, " It's type is: ", reflect.TypeOf(str))
	fmt.Println(f, " It's type is: ", reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(1.0))

	fmt.Println('A')
	fmt.Println(trueBoolean, falseBoolean)

	//quantity = 2
	//length, width = 10, 20
	//customerName = "Sandeep Gholve"

	fmt.Println(quantity, length, width, customerName, flag)
}
