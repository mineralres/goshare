package main

import (
	"log"

	"github.com/mineralres/goshare/pkg/httpapi"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var h httpapi.HTTPHandler
	h.Prepare("3030")
}
