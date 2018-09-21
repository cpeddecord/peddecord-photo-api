package makinjson

import "os"

var accessKey = os.Getenv("DO_SPACE_KEY")
var secKey = os.Getenv("DO_SPACE_SECRET")

var endpoint = "nyc3.digitaloceanspaces.com"
var project = "cpeddecord"
var bucket = "peddecord-photo"
