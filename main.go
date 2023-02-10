package main

import (
	"net/http"
    "fmt"
    "log"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("someBoxName", "./public")

    // Get the string representation of a file, or an error if it doesn't exist:
    // s, err := box.FindString("index.html")
    
    // Get the []byte representation of a file, or an error if it doesn't exist:
    s, err := box.Find("index.html")
    
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(s)


	http.Handle("/", http.FileServer(box))
	http.ListenAndServe(":3000", nil)
}