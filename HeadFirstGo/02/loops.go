package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	timeAndDate()
	stringReplacer()
}

func timeAndDate() {
	var currentTime time.Time = time.Now()
	hour, minute, seconds := currentTime.Clock()
	fmt.Println("Current time is: ", fmt.Sprintf("%d:%d:%d", hour, minute, seconds), " and year is: ", currentTime.Year())
}

func stringReplacer() {
	name := "Sandeep Gholve"
	replacer := strings.NewReplacer("e", "#")
	fmt.Println(replacer.Replace(name))
}
