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
	Path string `json:"path"`
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(response{
			commitHash,
			buildDate,
			r.URL.Path,
		})
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
	fmt.Println("Gophers at port 3000")
	http.ListenAndServe(":3000", nil)
}
