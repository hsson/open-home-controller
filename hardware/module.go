package hardware

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	typeToggle     = 0
	configFileName = "modules.json"
)

// Module represents a physical module connected to an Arduino
type Module struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Pin         int    `json:"pin"`
	Type        int    `json:"type"`
}

func initModules() ([]Module, error) {
	// Check if there is any config file
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		// File does not exist, create it
		var configFile *os.File
		configFile, err = os.Create(configFileName)
		defer configFile.Close()
		if err != nil {
			return nil, err
		}
		return []Module{}, nil
	}
	// File exists, open it
	jsonFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}
	var modules []Module
	err = json.Unmarshal(jsonFile, &modules)
	if err != nil {
		return nil, err
	}
	return modules, nil
}
