package sensor

import (
	"encoding/json"
	"net/http"

	"github.com/hsson/open-home-controller/hardware"
)

// Index lists all available modules
func Index(w http.ResponseWriter, r *http.Request) {
	modules := hardware.GetModules()
	enc := json.NewEncoder(w)
	err := enc.Encode(modules)
	if err != nil {
		http.Error(w, "Couldn't encode modules to JSON", http.StatusInternalServerError)
	}
}
