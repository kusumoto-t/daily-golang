package main

import "fmt"

func main() {

	var number1 int = 10
	var number2 int = 5

	result := addNumber(number1, number2)

	fmt.Println("Result of addition:", result)
}

func addNumber(a, b int) int {
	return a + b
}
