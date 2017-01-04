package main

import (
	"io/ioutil"
	"log"
	"strings"

	"time"

	"os"

	"github.com/tarm/serial"
)

const (
	baudRate    = 115200
	initTime    = time.Second * 2
	readTimeout = time.Second * 5
)

func main() {
	arduino := findArduino()
	if arduino == "" {
		log.Fatal("Couldn't find Arduino")
		os.Exit(1)
	}
	c := &serial.Config{Name: arduino, Baud: baudRate, ReadTimeout: readTimeout}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(initTime)

	n, err := s.Write([]byte("ping"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
}

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
