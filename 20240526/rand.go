package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < 30; i++ {
		FizzBuzz(rand.Intn(15))
	}

}

func FizzBuzz(input int) {
	if input%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if input%5 == 0 {
		fmt.Println("Fizz")
	} else if input%3 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(input)
	}

}
