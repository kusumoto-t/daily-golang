package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// BinarySearch
// ソートされた整数をスライスから指定された値で２分探索（バイナリサーチ）
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

	return -1 //見つからなかった場合
}

func main() {
	sortedArray := make([]int, 1000000)

	for i := range sortedArray {
		sortedArray[i] = i + 1
	}
	sort.Ints(sortedArray)

	// ランダムな探索対象値を生成
	rand.Seed(time.Now().UnixNano())

	target := rand.Intn(1000000) + 1

	startTime := time.Now() //開始時刻を記録

	result := BinarySearch(sortedArray, target)

	endTime := time.Now()                 //終了時刻を記録
	elapsedTime := endTime.Sub(startTime) //経過時間を計算

	if result != -1 {
		fmt.Printf("要素 %d は配列のインデックス %d にあります。\n", target, result)
	} else {
		fmt.Printf("要素 %d は配列に存在しません。\n", target)
	}

	fmt.Printf("処理時間： %v\n", elapsedTime)
}