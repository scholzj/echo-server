package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	address := ":9090"

	if len(os.Args) > 1 {
		address = fmt.Sprintf(":%s", os.Args[1])
	}

	server, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Failed to listen on address "+address, err)
		os.Exit(1)
	}

	defer server.Close()
	fmt.Printf("Listening on %v\n", server.Addr())

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Handling connection " + conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err == io.EOF {
				fmt.Println("Reached EOF")
			} else {
				fmt.Println("Failed to read bytes", err)
			}

			return
		}

		fmt.Printf("Echoing %v bytes\n", len(bytes))
		n, err := conn.Write(bytes)
		if err != nil {
			fmt.Println("Failed to echo bytes", err)
			return
		}

		fmt.Printf("Echoed %v bytes\n", n)
	}
}
