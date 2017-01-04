package sensor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hsson/open-home-controller/hardware"
)

// Index lists all available modules
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	modules := hardware.GetModules()
	enc := json.NewEncoder(w)
	err := enc.Encode(modules)
	if err != nil {
		http.Error(w, "Couldn't encode modules to JSON", http.StatusInternalServerError)
	}
}

// PostCommand sends a specified command to a specified module
func PostCommand(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	command := vars["command"]

	// This shouldn't happen since the mux only accepts numbers to this route
	if err != nil {
		http.Error(w, "Invalid Id, please try again.", http.StatusBadRequest)
		return
	}

	modules := hardware.GetModules()
	module, err := getModuleByID(modules, id)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	success := hardware.SendCommand(module.Pin, command)
	enc := json.NewEncoder(w)
	err = enc.Encode(successStruct{success})
}

// GetModuleValue gets the value of a specific module
func GetModuleValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	// This shouldn't happen since the mux only accepts numbers to this route
	if err != nil {
		http.Error(w, "Invalid Id, please try again.", http.StatusBadRequest)
		return
	}

	modules := hardware.GetModules()
	module, err := getModuleByID(modules, id)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	values := hardware.ReadStatus(module.Pin)
	var sens sensor
	sens.Name = module.Name
	sens.Description = module.Description
	sens.ID = module.ID
	sens.Pin = module.Pin
	sens.Type = module.Type
	sens.Values = values
	fmt.Println(values)
	enc := json.NewEncoder(w)
	err = enc.Encode(sens)
}
