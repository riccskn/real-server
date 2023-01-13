package handler

import (
	"bufio"
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	messages := make(chan string, 10)

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			messages <- scanner.Text()
		}
	}()

	go func() {
		for message := range messages {
			fmt.Fprintf(conn, "Message:: %s\n", message)
		}
	}()
}
