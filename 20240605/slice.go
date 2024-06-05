package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())

	input := rand.Intn(10)
	capacity := input + 5

	s := make([]int, input, capacity)

	fmt.Println(len(s))
	fmt.Println(cap(s))
	for i, v := range s {
		fmt.Println(i, ":", v)
	}
}