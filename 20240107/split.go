package main

import "fmt"

func main() {
	// 文字列
	str := "Hello Golang!"

	for _, char := range str {
		fmt.Printf("%c\n", char)
	}

}
