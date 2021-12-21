package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/sacOO7/gowebsocket"
)

type Message struct {
	Args []Node
}

type Node struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Properties []struct {
		Host  string `json:"host"`
		Port  int    `json:"port"`
		Slave int    `json:"slave"`
	} `json:"json"`
}

var message Message

func printWebSocket(scheme string, ip string, seconds int) {

	alive := true

	socket := gowebsocket.New(scheme + "://" + ip + "/ws/db")

	socket.OnConnected = func(socket gowebsocket.Socket) {
		//log.Println("Connected to server")
		//writeFile("Seconds: " + strconv.Itoa(seconds) + " Connected to server")
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		//log.Println("Recieved connect error ", err)
		writeFile("Seconds: " + strconv.Itoa(seconds) + " Fail")
		alive = false
	}

	socket.OnTextMessage = func(jsonmsg string, socket gowebsocket.Socket) {
		json.Unmarshal([]byte(jsonmsg), &message)
		/*for _, arg := range message.Args {
			log.Println(arg.Name)
			writeFile("Node name: " + arg.Name)
		}*/
		writeFile("Seconds: " + strconv.Itoa(seconds) + " Pass")
		alive = false
		socket.Close()
	}

	socket.OnBinaryMessage = func(data []byte, socket gowebsocket.Socket) {
		log.Println("Recieved binary data ", data)
	}

	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Recieved ping " + data)
	}

	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Recieved pong " + data)
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		//log.Println("Disconnected from server ")
		//writeFile("Disconnected from server ")
	}

	socket.Connect()

readLoop:
	for {
		if !alive {
			break readLoop
		}
	}

}
