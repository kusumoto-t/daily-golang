package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/api/", apiHandler)

	port := 8080

	fmt.Printf("Starting Server port :%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Println("Error Starting Server!")
	}

}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API Endpoint")
}
