package main

import "fmt"

func isPalindrome(input string) bool {
	length := len(input)

	for i := 0; i < length/2; i++ {
		if input[i] != input[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var userInput string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&userInput)

	if isPalindrome(userInput) {
		fmt.Println("これは palindromeです")
	} else {
		fmt.Println("これは palindromeではありません")
	}

}
