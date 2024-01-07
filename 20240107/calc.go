package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 int = 5
	
	resultAdd := add(num1, num2)
	fmt.Println("Result of addition:", resultAdd)

	resultSubtract := subtract(num1, num2)
	fmt.Println("Result of subtract:", resultSubtract)

	resultMultiply := multiply(num1, num2)
	fmt.Println("Result of multiply:", resultMultiply)

	resultDivide := divide(num1, num2)
	fmt.Println("Result of divide:", resultDivide)
}

// 足し算
func add(a, b int) int {
	return a + b
}

// 引き算
func subtract(a, b int) int {
	return a - b
}

// 掛け算
func multiply(a, b int) int {
	return a * b
}

// 割り算
func divide(a, b int) int {
	if b == 0 {
		fmt.Println("Error: division by zero")
		return 0
	}
	return a / b
}
