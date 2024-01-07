package main

import "fmt"

func main() {
	str := "Hello World!"

	for i, char := range str {
		fmt.Printf("%d : %c \n", i,char)
	}
}
