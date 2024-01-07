package main

import "fmt"
import "time"

func main() {
	now := time.Now()
	// 20060102と書くのがGolangの決まり。アメリカ式の時間順は1月2日午後3時4分5秒2006年となっているため
	fmt.Println("Hello World!" + now.Format("20060102"))
}
