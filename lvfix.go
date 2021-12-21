package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/samuelventura/go-modbus"
	"github.com/samuelventura/go-serial"
)

func powerOn(secondsOn int, master modbus.Master) {
	for i := 0; i < secondsOn; i++ {
		err := master.WriteDo(1, 4096, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Println("")
}

func powerOff(master modbus.Master) {
	err := master.WriteDo(1, 4096, false)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)
}

//(cd sample; go run .)
func main() {
	scheme := ""
	ip := ""
	if len(os.Args) > 2 {
		scheme = os.Args[1]
		ip = os.Args[2]
	}
	const MODULO = 32
	log.SetFlags(log.Lmicroseconds)
	mode := &serial.Mode{}
	mode.BaudRate = 9600
	mode.DataBits = 8
	mode.Parity = serial.NoParity
	mode.StopBits = serial.OneStopBit
	trans, err := serial.NewSerialTransport("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal(err)
	}
	defer trans.Close()
	trans.DiscardOn()
	modbus.EnableTrace(false)
	master := modbus.NewRtuMaster(trans, 400)

	counter := 1
	for {
		secondsOn := counter % MODULO
		fmt.Printf("Seconds: %v", secondsOn)
		powerOn(secondsOn, master)
		if secondsOn > 12 {
			printWebSocket(scheme, ip, secondsOn)
		}
		powerOff(master)
		counter++
	}
}
