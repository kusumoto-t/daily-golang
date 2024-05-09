package main

import "fmt"

func main() {
	var numA int = 10
	var numB int = 20

	fmt.Println("Add:", add(numA, numB))
	fmt.Println("Substract:", substract(numA, numB))
	fmt.Println("Multiply:", multiply(numA, numB))
	fmt.Println("Devide:", devide(numA, numB))
}

func add(a, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func devide(a, b int) int {
	return a / b
}
