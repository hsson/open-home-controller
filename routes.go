package main

import (
	"net/http"

	"github.com/hsson/open-home-controller/modules/sensor"
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
	// Hardware modules
	Route{"IndexSensors", "GET", "/api/sensors", false, sensor.Index},
	Route{"GetModuleValue", "GET", "/api/sensors/{id:[0-9]+}", false, sensor.GetModuleValue},
	Route{"PostCommand", "POST", "/api/sensors/{id:[0-9]+}/{command:[a-z]}", false, sensor.PostCommand},
}
