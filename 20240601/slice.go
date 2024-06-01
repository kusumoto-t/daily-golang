package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)

	for i, v := range a {
		fmt.Println(i+1, "番目:", v)
	}
}