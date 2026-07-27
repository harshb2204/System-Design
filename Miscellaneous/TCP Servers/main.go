package main

import (
	"log"
	"net"
	"time"
)

func do(connection net.Conn) {

	var buf []byte = make([]byte, 1024)

	_, err := connection.Read(buf)

	if err != nil {
		log.Fatal(err)
	}
	// mimicing a process that takes time to complete

	log.Println("Doing some processing")

	time.Sleep(8 * time.Second)

	connection.Write([]byte(
		"HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/plain\r\n" +
			"\r\n" +
			"Hello from server",
	))
	connection.Close()
}

func main() {

	listener, err := net.Listen("tcp", ":1729") // listen for tcp protocol on port 1729

	if err != nil {
		log.Fatal(err)
	}

	for {

		log.Println("Waiting for client to connect")

		connection, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Client connected")

		go do(connection)

	}

}
