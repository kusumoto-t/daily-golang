package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api/", apiHandler)

	port := 8080

	fmt.Printf("Server start port %d ... \n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Println("Error on Starting Server")
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Golang!")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello api!")
}
