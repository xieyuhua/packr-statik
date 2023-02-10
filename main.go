package main

import (
	"net/http"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("someBoxName", "./public")

	http.Handle("/", http.FileServer(box))
	http.ListenAndServe(":3000", nil)
}