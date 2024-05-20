package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Seed(time.Now().UnixNao()) // こちらは1.20から非推奨になった
	rand.NewSource(time.Now().UnixNano())

	fmt.Println(rand.Intn(100))
}
