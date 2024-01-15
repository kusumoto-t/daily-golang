package main

import "fmt"

func main() {
	// 整数型のスライスを作成（初期長さ：0、初期容量：5）
	stack := make([]int, 0, 5) // 初期容量は任意

	// スタックに要素を追加
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)

	fmt.Println("スタックの内容:", stack)
}
