package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	service := ":8080"

	fmt.Println("Server starting..")

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	fmt.Printf("Server listening at %s\n", tcpAddr)

	listener, err := net.ListenTCP("tcp", tcpAddr)

	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go HandleConnection(conn)
	}

}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	buf := make([]byte, 1024)

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "Error reading message: %s\n", err.Error())
				break
			}
		}

		writer.Write(buf[:n])
		writer.Flush()
	}
}

func checkError(err error) {

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
