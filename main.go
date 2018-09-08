package main

import (
	"fmt"
	"net/http"
)

// look at these sweet comments
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "asuh")
	})

	http.ListenAndServe(":8080", nil)
}
