package main

import "fmt"


func main() {
	var number1 int = 10
	var number2 int = 20

	result := addNumbers(number1, number2)

	fmt.Println("Result of addition:", result)
}

func addNumbers(a, b int) int {
	return a + b
}
