//go:build !devel

// Command test_server is a simple web server that serves the frontend/dist.
package main

import (
	"embed"
	"flag"
	"log"
	"net/http"

	"github.com/rusq/httpex"
)

//go:embed frontend/dist
var fsys embed.FS

func main() {
	flag.Parse()

	spa, err := httpex.NewVueSPA(fsys, "frontend/dist", "assets")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", spa) // main application
	http.ListenAndServe(*listenAddr, nil)
}
