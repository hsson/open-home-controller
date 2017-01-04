package main

import (
	"net/http"

	"github.com/hsson/hardware"
)

func main() {
	r := NewRouter()
	hardware.Initialize()
	http.Handle("/", r)
}
