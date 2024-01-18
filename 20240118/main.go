package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ハンドラを登録
	http.HandleFunc("/", helloHandler)

	// HTTPサーバ起動
	port := 8080
	fmt.Printf("Server is listen on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error stating server:", err)
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Golang!")
}
