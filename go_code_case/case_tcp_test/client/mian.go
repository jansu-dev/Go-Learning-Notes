package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// connect with server
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:2000 failed err: ", err)
		return
	}
	sendData(conn)
}

func sendData(conn net.Conn) {
	// send data
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please say: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}

	defer conn.Close()
}
