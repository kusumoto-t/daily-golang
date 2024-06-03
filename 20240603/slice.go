package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())

	rLength := rand.Intn(10)

	capacity := rLength + 10
	array := make([]int, rLength, capacity)

	fmt.Printf("Initial length -> %d\n", len(array))
	fmt.Printf("Initial capacity -> %d\n", cap(array))

	for i, v := range array {
		fmt.Println(i, ":", v)
	}
}