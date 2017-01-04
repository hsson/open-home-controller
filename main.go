package main

import (
	"bufio"
	"fmt"
	"log"

	"time"

	"os"

	"strconv"

	"github.com/tarm/serial"
)

const (
	baudRate     = 115200
	initTime     = time.Second * 2
	readTimeout  = time.Second * 5
	statusAction = "s"
	trueLiteral  = "True"
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
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter pin: ")
		pin, _ := reader.ReadString('\n')
		pinInt, _ := strconv.Atoi(pin[:len(pin)-1])
		fmt.Println(pinInt)
		fmt.Print("Enter action: ")
		action, _ := reader.ReadString('\n')
		success := sendCommand(uint16(pinInt), action[:1])
		log.Println("Success:", success)
	}
}

func sendCommand(pin uint16, action string) bool {
	command := Command{pin, action}
	log.Println("Sending:", command.parse())
	n, err := serialPort.Write(command.parseBytes())

	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = serialPort.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	res := string(buf[:n])
	return res == trueLiteral
}
