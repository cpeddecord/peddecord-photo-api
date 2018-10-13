package main

import (
	"github.com/graphql-go/graphql"
)

type Image struct {
	ID          string
	Title       string
	CreatedDate string
	URL         string
	ImageHeight string
	ImageWidth  string
}

type Gallery struct {
	ID        int
	Category  string
	Slug      string
	Lede      string
	SubLede   string
	Caption   string
	Thumbnail Image
	Images    []Image
}

type Galleries []Gallery

var ImageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Image",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"createdDate": &graphql.Field{
			Type: graphql.String,
		},
		"url": &graphql.Field{
			Type: graphql.String,
		},
		"imageHeight": &graphql.Field{
			Type:        graphql.Int,
			Description: "Pixel Height",
		},
		"imageWidth": &graphql.Field{
			Type:        graphql.Int,
			Description: "Pixel Width",
		},
	},
})

var GalleryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Gallery",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"category": &graphql.Field{
			Type: graphql.String,
		},
		"slug": &graphql.Field{
			Type: graphql.String,
		},
		"lede": &graphql.Field{
			Type: graphql.String,
		},
		"sublede": &graphql.Field{
			Type: graphql.String,
		},
		"caption": &graphql.Field{
			Type: graphql.String,
		},
		"thumbnail": &graphql.Field{
			Type: ImageType,
		},
		"images": &graphql.Field{
			Type: graphql.NewList(ImageType),
		},
	},
})
