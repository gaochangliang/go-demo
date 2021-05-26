package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	g := gee.New()
	g.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	g.Run(":8080")
}
