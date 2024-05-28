package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())

	var fizzBuzzCnt = 0
	var fizzCnt = 0
	var buzzCnt = 0

	for i := 0; i < 30; i++ {
		num := rand.Intn(15)

		if num%15 == 0 {
			fizzBuzzCnt++
			fmt.Println("FizzBuzz")
		} else if num%5 == 0 {
			fizzCnt++
			fmt.Println("Fizz")
		} else if num%3 == 0 {
			buzzCnt++
			fmt.Println("Buzz")
		} else {
			fmt.Println(num)
		}
	}
	fmt.Println("=========================")
	fmt.Println("fizzBuzzCnt:", fizzBuzzCnt)
	fmt.Println("fizzCnt:", fizzCnt)
	fmt.Println("buzzCnt:", buzzCnt)
}