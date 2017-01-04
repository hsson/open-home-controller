package hardware

import (
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/tarm/serial"
)

const (
	baudRate     = 115200
	bufferSize   = 512
	initTime     = time.Second * 2
	readTimeout  = time.Second * 5
	statusAction = "s"
	trueLiteral  = "True"
)

var (
	serialPort *serial.Port
	modules    []Module
	mutex      = &sync.Mutex{}
)

// Initialize sets up the connection to the hardware and reads the config
func Initialize() {
	// Init connection to Arduino
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

	// Read config
	log.Println("Initializing modules...")
	modules, err = initModules()
	if err != nil {
		log.Fatal("Could not read config\n", err)
		os.Exit(1)
	}
	log.Println("Module config loaded")

	// DONE
	log.Println("Initialization complete")

}

// ReadStatus gets the value of a given hardware module
func ReadStatus(pin int) []string {
	mutex.Lock()
	defer mutex.Unlock()
	command := Command{pin, statusAction}
	log.Println("Reading status from pin:", pin)
	n, err := serialPort.Write(command.parseBytes())
	if err != nil {
		log.Fatalf("Couldn't write to pin %v, error: %v\n", pin, err)
	}

	buf := make([]byte, bufferSize)
	n, err = serialPort.Read(buf)
	if err != nil {
		log.Fatalf("Couldn't read status from pin %v, error: %v\n", pin, err)
	}
	res := string(buf[:n])
	return strings.Split(res, ";")
}

// SendCommand sends an action to a given hardware modulue
func SendCommand(pin int, action string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	command := Command{pin, action}
	log.Println("Sending:", command.parse())
	n, err := serialPort.Write(command.parseBytes())

	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, bufferSize)
	n, err = serialPort.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	res := string(buf[:n])
	return res == trueLiteral
}

// GetModules returns the available hardware modules
func GetModules() []Module {
	return modules
}
