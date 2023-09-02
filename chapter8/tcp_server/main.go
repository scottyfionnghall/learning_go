package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// listen on a port
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accpet a connection
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(connection)
	}
}

func handleServerConnection(connection net.Conn) {
	// recevie the message
	var message string
	err := gob.NewDecoder(connection).Decode(&message)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Recived", message)
	}
	connection.Close()
}

func client() {
	//connect to the cerver
	connection, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	// send the message
	message := "Hello, World"
	fmt.Println("Sending", message)
	err = gob.NewEncoder(connection).Encode(message)
	if err != nil {
		fmt.Println(err)
	}

	connection.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
