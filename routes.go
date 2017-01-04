package main

import (
	"net/http"

	"github.com/hsson/open-home-controller/modules/modules/sensor"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Secured     bool
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	// Projects
	Route{"IndexSensors", "GET", "/api/sensors", false, sensor.Index},
}
