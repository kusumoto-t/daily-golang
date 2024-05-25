package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		num := 1 + rand.Intn(10-1) // non zero
		//		num := rand.Intn(10)       // include zero
		if num == 0 {
			fmt.Println("rand num is zero")
		} else {
			fmt.Println(num)
		}
	}
}
