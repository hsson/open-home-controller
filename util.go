package main

import (
	"io/ioutil"
	"strings"
)

// Finds the serial port for the Arduino
func findArduino() string {
	contents, _ := ioutil.ReadDir("/dev")
	for _, f := range contents {
		if strings.Contains(f.Name(), "tty.usbserial") || strings.Contains(f.Name(), "ttyACM") {
			return "/dev/" + f.Name()
		}
	}
	return ""
}
