package main

import (
	"fmt"
	"net/http"
)

// look at these sweet comments
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "yo")
	})

	fmt.Println("starting on port 8080")
	http.ListenAndServe(":8080", nil)
}
