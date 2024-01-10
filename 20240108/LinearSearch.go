package main

import (
	"fmt"
	"time"
)

// LinearSearch は整数のスライスから指定された値を線形探索で検索する関数です。
func LinearSearch(arr []int, target int) int {
	for i, element := range arr {
		if element == target {
			return i // 見つかった場合、要素のインデックスを返す
		}
	}
	return -1
}

func main() {
	// 整数のスライス作成
	slice := []int{4, 2, 7, 1, 9, 5, 3, 8, 6}

	target := 6

	startTime := time.Now() // 開始時刻を記録

	result := LinearSearch(slice, target)

	endTime := time.Now() // 終了時刻を記録
	elapsedTime := endTime.Sub(startTime)

	if result != -1 {
		fmt.Printf("要素 %d はスライスのインデックス %d にあります。\n", target, result)
	} else {
		fmt.Printf("要素 %d はスライスに存在しません。", target)
	}

	fmt.Printf("処理時間： %v\n", elapsedTime)
}