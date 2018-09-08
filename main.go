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

	fmt.Println("going to start now, k?")
	http.ListenAndServe(":80", nil)
}
