package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ハンドラを登録
	http.HandleFunc("/", helloHandler)

	// HTTPサーバを起動
	port := 8080
	fmt.Printf("Server is listening on : %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// レスポンスを書き込む
	fmt.Fprint(w, "Hello, World!")
}
