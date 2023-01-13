package main

import (
	"fmt"
	"net"
)
import "server/handler"

func main() {

	listener, err := net.Listen("tcp", ":8080")

	if err == nil {
		fmt.Print("Starting server at port: 8080")
	} else {
		fmt.Println("Couldn't start the server at port :8080")
	}

	for {

		conn, _ := listener.Accept()

		go handler.HandleConnection(conn)

	}

}
