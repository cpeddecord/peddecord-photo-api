package main

import (
	"github.com/graphql-go/graphql"
)

var galleryField = &graphql.Field{
	Type:        graphql.NewList(GalleryType),
	Description: "Get one or multiple galleries",
	Args: graphql.FieldConfigArgument{
		"slug": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		slugQuery, slugIsOK := params.Args["slug"].(string)

		var galleries Galleries

		if slugIsOK {
			for _, gallery := range MasterGalleryList {
				if gallery.Slug == slugQuery {
					galleries = append(galleries, gallery)
				}
			}

			return galleries, nil
		}

		return MasterGalleryList, nil
	},
}

var imageField = &graphql.Field{
	Type:        ImageType,
	Description: "Get single image by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		idQuery, idIsOK := params.Args["id"].(string)

		if idIsOK {
			for _, gallery := range MasterGalleryList {
				for _, img := range gallery.Images {
					if img.ID == idQuery {
						return img, nil
					}
				}
			}
		}

		return nil, nil
	},
}

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"galleries": galleryField,
		"image":     imageField,
	},
})
