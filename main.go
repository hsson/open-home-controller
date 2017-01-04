package main

import (
	"net/http"

	"github.com/hsson/open-home-controller/hardware"
)

func main() {
	r := NewRouter()
	hardware.Initialize()
	http.Handle("/", r)
}
