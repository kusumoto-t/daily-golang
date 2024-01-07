package main

import "fmt"

func main() {
	// 1から10までの数字を表示する
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("=============")

	// 10から1までの数を表示するループ
	for j := 10; j >= 1; j-- {
		fmt.Println(j)
	}
}