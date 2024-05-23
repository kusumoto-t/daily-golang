package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano()) // rand.SeedはDeperecated

	for i := 0; i < 30; i++ {
		fmt.Println("num is ", 1+rand.Intn(10-1)) // 1 + にすることで0は出力されない
	}
}