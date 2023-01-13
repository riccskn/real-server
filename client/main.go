package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", ":8080")

	messages := make(chan string, 10)

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			messages <- scanner.Text()
		}
	}()

	go func() {
		for message := range messages {
			fmt.Println(message)
		}
	}()

	writer := bufio.NewWriter(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		writer.WriteString(scanner.Text() + "\n")
		writer.Flush()
	}
}
