package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var MasterGalleryList []Gallery

func init() {
	// TODO: populate with proper data, maybe do some injection into a closure, I dunno...
	MasterGalleryList = append(MasterGalleryList, FixtureGallery1, FixtureGallery2)
}

func main() {
	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		basePath = "/"
	}
	queryPath := basePath + "graphql"

	var schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: RootQuery,
	})

	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	http.HandleFunc(basePath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "bread")
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🍞")
	})

	http.Handle(queryPath, h)

	fmt.Println("Gophers at port 3000")
	http.ListenAndServe(":3000", nil)
}
