package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	minio "github.com/minio/minio-go"
)

type response struct {
	Hash string `json:"hash"`
	Date string `json:"date"`
}

var commitHash = os.Getenv("COMMIT_REF")
var buildDate = os.Getenv("BUILD_DATE")

var accessKey = os.Getenv("DO_SPACE_KEY")
var secKey = os.Getenv("DO_SPACE_SECRET")

var endpoint = "nyc3.digitaloceanspaces.com"
var project = "cpeddecord"
var bucket = "peddecord-photo"

// look at these sweet comments
func main() {
	client, err := minio.New(endpoint, accessKey, secKey, true)
	if err != nil {
		log.Fatal(err)
	}

	res := response{
		commitHash,
		buildDate,
	}

	data, _ := json.Marshal(res)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var s []string
		done := make(chan struct{})
		for o := range client.ListObjects(bucket, "", true, done) {
			if o.Err != nil {
				log.Println(o.Err)
				return
			}

			s = append(s, o.Key)
		}

		data, _ := json.Marshal(s)
		w.Write(data)
	})

	fmt.Println("Build Hash: ", commitHash)
	fmt.Println("Build Date: ", buildDate)
	fmt.Println("starting on port 8080")
	http.ListenAndServe(":8080", nil)
}
