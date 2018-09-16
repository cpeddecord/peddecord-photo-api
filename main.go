package main

import (
	"fmt"
	"net/http"
	"os"
)

var commitHash = os.Getenv("COMMIT_REF")
var buildDate = os.Getenv("BUILD_DATE")

// look at these sweet comments
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "seeing if these pipes truly work\n\nbuild %s at %s", commitHash, buildDate)
	})

	fmt.Println("Build Hash: ", commitHash)
	fmt.Println("Build Date: ", buildDate)
	fmt.Println("starting on port 8080")
	http.ListenAndServe(":8080", nil)
}
