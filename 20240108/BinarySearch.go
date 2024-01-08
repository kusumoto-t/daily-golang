package main

import (
	"fmt"
	"time"
)

// BinarySearch はソートされた整数のスライスから指定された値をバイナリサーチで検索する関数です。
func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		if guess == target {
			return mid
		} else if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1 // 見つからなかった場合
}

func main() {
	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 1

	startTime := time.Now() // 開始時刻を記録

	result := BinarySearch(sortedArray, target)

	endTime := time.Now()                 // 終了時刻を記録
	elapsedTime := endTime.Sub(startTime) // 経過時間を計算

	if result != -1 {
		fmt.Printf("要素 %d は配列インデックス %d にあります。\n", target, result)
	} else {
		fmt.Printf("要素 %d は配列に存在しません。\n", target)
	}

	fmt.Printf("処理時間： %v\n", elapsedTime)
}