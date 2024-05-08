package main

import "fmt"

func main() {

	var numA = 10
	var numB = 5

	fmt.Println("Result of addition:", add(numA, numB))
	fmt.Println("Result of substract:", substract(numA, numB))
	fmt.Println("Result of multiply:", multiply(numA, numB))
	fmt.Println("Result of devide:", devide(numA, numB))
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
