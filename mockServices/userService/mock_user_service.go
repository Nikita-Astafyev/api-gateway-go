package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Response from User-Service (Mock)")
	})

	log.Println("Mock User-Service running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
