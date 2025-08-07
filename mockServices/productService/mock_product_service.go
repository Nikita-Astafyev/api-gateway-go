package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Response from Product-Service (Mock)")
	})

	log.Println("Mock Product-Service running on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
