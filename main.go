package main

import (
	"log"
	"net/http"

	"github.com/hsson/open-home-controller/hardware"
)

func main() {
	r := NewRouter()
	hardware.Initialize()
	http.Handle("/", r)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
