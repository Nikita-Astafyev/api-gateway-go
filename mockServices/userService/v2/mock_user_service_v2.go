package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Response from User-Service V2 (Mock)")
	})

	log.Println("Mock User-Service V2 running on :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
