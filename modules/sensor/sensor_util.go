package sensor

import (
	"errors"

	"github.com/hsson/open-home-controller/hardware"
)

func getModuleByID(modules []hardware.Module, id int) (hardware.Module, error) {
	for _, module := range modules {
		if module.ID == id {
			return module, nil
		}
	}
	return hardware.Module{}, errors.New("Module with specified ID not found")
}
