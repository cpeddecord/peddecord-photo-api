package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type response struct {
	Hash   string `json:hash`
	Date   string `json:date`
	Secret string `json:secret`
}

var commitHash = os.Getenv("COMMIT_REF")
var buildDate = os.Getenv("BUILD_DATE")
var secret = os.Getenv("TEST_SECRET")

// look at these sweet comments
func main() {
	res := response{
		commitHash,
		buildDate,
		secret,
	}

	data, _ := json.Marshal(res)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	fmt.Println("Build Hash: ", commitHash)
	fmt.Println("Build Date: ", buildDate)
	fmt.Println("starting on port 8080")
	http.ListenAndServe(":8080", nil)
}
