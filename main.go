package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "docker-api-service running on port 8005")
	})

	fmt.Println("Server starting on :8005")
	err := http.ListenAndServe(":8005", nil)
	if err != nil {
		panic(err)
	}
}
