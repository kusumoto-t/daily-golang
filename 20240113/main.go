package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func IsPrime(num int) bool {
	if num < 2 {
		return false // 2未満の数は素数ではない
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false // 割り切れたら素数ではない
		}
	}

	return true
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("判定したい数を入力してください")
		return
	}

	number, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("数値を入力してください")
		return
	}

	if IsPrime(number) {
		fmt.Printf("%d は素数です。\n", number)
	} else {
		fmt.Printf("%d は素数ではありません。\n", number)
	}
}
