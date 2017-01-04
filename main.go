package main

import (
	"log"

	"time"

	"os"

	"github.com/tarm/serial"
)

const (
	baudRate    = 115200
	initTime    = time.Second * 2
	readTimeout = time.Second * 5
)

var (
	serialPort *serial.Port
)

func init() {
	log.Println("Finding Arduino...")
	arduino := findArduino()
	if arduino == "" {
		log.Fatalln("Could not find Arduino")
		os.Exit(1)
	}
	log.Println("Arduino found on:", arduino)

	log.Println("Initializing port...")
	config := &serial.Config{Name: arduino, Baud: baudRate, ReadTimeout: readTimeout}
	s, err := serial.OpenPort(config)
	if err != nil {
		log.Fatalln("Could not open port!\n", err)
	}
	serialPort = s
	log.Println("Port opened")
	time.Sleep(initTime)
	log.Println("Initialization complete")

}

func main() {
	n, err := serialPort.Write([]byte("ping"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = serialPort.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
}
