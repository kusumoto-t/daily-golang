package main

import (
	"fmt"
	"strconv"
)

func main() {
	numA := 1
	numB := 2
	fmt.Println("Hello Golang!")
	s := strconv.Itoa(numA + numB)
	fmt.Println("Sum is " + s)
}
