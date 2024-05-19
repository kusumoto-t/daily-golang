package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 乱数のシードを現在の時間に基づいて初期化
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Rand number is", rand.Intn(100))
}